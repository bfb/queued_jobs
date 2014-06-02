package qj

import (
	"encoding/json"
)

type Job struct {
	Jid      string        `json:"jid"`
	Queue    string        `json:"queue"`
	Args     []interface{} `json:"args"`
	PostedAt int64         `json:"posted_at"`
}

func JobMarshal(job *Job) string {
	res, _ := json.Marshal(job)
	return string(res)
}

func JobUnmarshal(job string) (*Job, error) {
	res := &Job{}
	json.Unmarshal([]byte(job), &res)
	return res, nil
}
