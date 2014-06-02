package qj

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

type settings struct {
	namespace string
	Pool      *redis.Pool
}

var Settings *settings
var managers = make(map[string]*manager)
var Logger = log.New(os.Stdout, "=> ", log.Ldate|log.Lmicroseconds)

func Setup(uri string, namespace string, size int) {
	Settings = &settings{
		namespace,
		NewRedisPool(uri, size),
	}
}

func NewRedisPool(uri string, size int) *redis.Pool {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	var password string = ""

	if u.User != nil {
		password, _ = u.User.Password()
	}

	var db string = ""
	if u.Path != "" {
		db = strings.Replace(u.Path, "/", "", 1)
	}

	return &redis.Pool{
		MaxIdle:     size,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", u.Host)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if db != "" {
				if _, err := c.Do("SELECT", db); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func KeyName(key string) string {
	return Settings.namespace + ":" + key
}

func Workers(queue string, task Task, concurrency int) {
	managers[queue] = NewManager(queue, task, concurrency)
}

func Start() {
	for _, manager := range managers {
		manager.start()
	}

	WaitWorkers()
}
