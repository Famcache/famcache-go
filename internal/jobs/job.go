package jobs

import "github.com/famcache/famcache-go/domain"

type Job struct {
	requestID  string
	name       string
	jobID      string
	delay      uint64
	periodic   bool
	listenChan chan struct{}
}

func NewJob(requestID, name string, delay uint64, periodic bool) domain.Job {
	return &Job{
		requestID:  requestID,
		name:       name,
		jobID:      "",
		delay:      delay,
		periodic:   periodic,
		listenChan: make(chan struct{}),
	}
}
