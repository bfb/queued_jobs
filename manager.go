package qj

import (
	"github.com/garyburd/redigo/redis"
	"sync"
)

type manager struct {
	queue       string
	concurrency int
	task        Task
	jobs        chan *Job
	workers     []*worker
	*sync.WaitGroup
}

func NewManager(queue string, task Task, concurrency int) *manager {
	m := &manager{
		queue,
		concurrency,
		task,
		make(chan *Job),
		make([]*worker, concurrency),
		&sync.WaitGroup{},
	}
	return m
}

func (m *manager) start() {
	m.Add(1)

	for i := 0; i < m.concurrency; i++ {
		m.workers[i] = NewWorker(m, i)
		Logger.Println("Starting worker", i, "on", m.queue, "queue.")
		go m.workers[i].start()
	}

	m.getJob()
}

func WaitWorkers() {
	for _, manager := range managers {
		manager.Wait()
	}
}

func (m *manager) getJob() {
	c := Settings.Pool.Get()
	defer c.Close()

	for {
		json_job, _ := redis.String(c.Do("brpoplpush", KeyName(m.queue), KeyName("processing"), 1))
		if json_job != "" {
			job, _ := JobUnmarshal(json_job)
			m.jobs <- job
		}
	}
}

func (m *manager) failedJob(job *Job, err error, trace string) {
	c := Settings.Pool.Get()
	defer c.Close()

	failure := &Failure{job.Jid, job.Queue, job.Args, job.PostedAt, err.Error(), trace}
	c.Do("lpush", KeyName("failures"), FailureMarshal(failure))
	c.Do("incr", KeyName("stats:failed"))
	c.Do("lrem", KeyName("processing"), 1, JobMarshal(job))
}

func (m *manager) doneJob(job *Job) {
	c := Settings.Pool.Get()
	defer c.Close()

	c.Do("incr", KeyName("stats:processed"))
	c.Do("lrem", KeyName("processing"), 1, JobMarshal(job))
}
