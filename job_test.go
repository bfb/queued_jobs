package qj

import (
	"testing"
)

func TestJobUnmarshal(t *testing.T) {
	json := "{\"queue\":\"high\",\"args\":[\"http://wikipedia.org\",10],\"posted_at\":1401231507,\"jid\":\"65a27cfe318d84a5ad94426b05e428b7e5209cd3\"}"
	job, _ := JobUnmarshal(json)

	if job.Queue != "high" {
		t.Error("Expected high, got ", job.Queue)
	}

	if job.PostedAt != 1401231507 {
		t.Error("Expected 1401231507, got ", job.PostedAt)
	}

	if job.Jid != "65a27cfe318d84a5ad94426b05e428b7e5209cd3" {
		t.Error("Expected 65a27cfe318d84a5ad94426b05e428b7e5209cd3, got ", job.Jid)
	}
}

func TestJobMarshal(t *testing.T) {
	json := "{\"queue\":\"high\",\"args\":[\"http://wikipedia.org\",10],\"posted_at\":1401231507,\"jid\":\"65a27cfe318d84a5ad94426b05e428b7e5209cd3\"}"
	job := Job{"high", Interface{"http://wikipedia.org", 10}, 1401231507, "65a27cfe318d84a5ad94426b05e428b7e5209cd3"}
	res, _ := JobMarshal(job)

	if res != json {
		t.Error("Expected json, got ", res)
	}
}
