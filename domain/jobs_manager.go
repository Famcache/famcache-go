package domain

type JobEvent interface {
	IsRegistered() bool
	IsExecute() bool
}

type JobReply interface {
	JobID() string
	RequestID() string
	Event() JobEvent
}

type JobsManager interface {
	Registry(job Job) Job
	Cancel(jobID string) error
}
