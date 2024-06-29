package domain

type Job interface {
	RequestID() string
	Name() string
	JobID() string
	Delay() uint64
	IsPeriodic() bool
	Listen() <-chan struct{}
	IsRegistered() bool
}
