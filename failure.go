package qj

import (
	"encoding/json"
)

type Failure struct {
	Jid       string        `json:"jid"`
	Queue     string        `json:"queue"`
	Args      []interface{} `json:"args"`
	PostedAt  int64         `json:"posted_at"`
	Error     string        `json:"error"`
	Backtrace string        `json:"backtrace"`
}

func FailureMarshal(failure *Failure) string {
	res, _ := json.Marshal(failure)
	return string(res)
}

func FailureUnmarshal(failure string) (*Failure, error) {
	res := &Failure{}
	json.Unmarshal([]byte(failure), &res)
	return res, nil
}
