package qj

import (
	"strconv"
	"testing"
)

func TestSetup(t *testing.T) {
	SetupTest()
	if Settings.Pool == nil {
		t.Error("Invalid Redis server")
	}

	if Settings.namespace != "production" {
		t.Error("Expected namespace production, got ", Settings.namespace)
	}
}

func TestNewRedisPool(t *testing.T) {
	SetupTest()

	if Settings.Pool.Get() == nil {
		t.Error("Expected successfull connection")
	}
}

func TestKeyName(t *testing.T) {
	key := KeyName("processing")

	if key != "production:processing" {
		t.Error("Expected production:processing, got", key)
	}
}

func TestConfigWorker(t *testing.T) {
	SetupTest()
	configWorker("high", nil, 20)
	m := managers["high"]

	if m.queue != "high" {
		t.Error("Expected high as queue, got ", m.queue)
	}

	if m.task != nil {
		t.Error("Expected nil task, got ", m.task)
	}
}

func TestStart(t *testing.T) {
	SetupTest()
	configWorker("high", nil, 20)
	total_managers := len(managers)

	if total_managers != 1 {
		t.Error("Expected 1 manager, got ", strconv.Itoa(total_managers))
	}
}

func SetupTest() {
	Setup("redis://localhost:6379", "production", 10)
}
