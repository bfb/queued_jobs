package qj

import (
	"runtime/debug"
)

type worker struct {
	manager   *manager
	job       *Job
	startedAt int64
	wid       int
}

func (w *worker) start() {
	// access jobs channel on manager
	for {
		defer func() {
			if r := recover(); r != nil {
				Logger.Println("Job", w.job.Jid, "failed:", r.(error), ".")
				w.failed(r.(error), debug.Stack())
				w.start()
			}
		}()

		select {
		case w.job = <-w.manager.jobs:
			if w.process() {
				w.done()
			}
		}
	}
}

func (w *worker) done() {
	Logger.Println("Job", w.job.Jid, "done.")
	w.manager.doneJob(w.job)
}

func (w *worker) process() bool {
	Logger.Println("Job", w.job.Jid, "in process.")
	w.manager.task(w.job.Args...)
	return true
}

func (w *worker) failed(err error, trace []byte) {
	w.manager.failedJob(w.job, err, string(trace))
}

func NewWorker(m *manager, wid int) *worker {
	return &worker{m, nil, 0, wid}
}
