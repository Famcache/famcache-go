package domain

type Status int

const (
	OK Status = iota
	Err
	Pending
)

type Task interface {
	Execute(data string, status Status)
	Wait()
	GetResult() (data string, status Status)

	IsSuccess() bool
	IsError() bool
	IsPending() bool
}
