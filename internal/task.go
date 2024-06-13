package internal

import (
	"sync"

	"github.com/famcache/famcache-go/domain"
)

type Task struct {
	wg *sync.WaitGroup

	data   string
	status domain.Status
}

func (t *Task) Execute(data string, status domain.Status) {
	t.data = data
	t.status = status

	t.wg.Done()
}

func (t *Task) Wait() {
	t.wg.Wait()
}

func (t *Task) GetResult() (data string, status domain.Status) {
	return t.data, t.status
}

func (t *Task) IsSuccess() bool {
	return t.status == domain.OK
}

func (t *Task) IsError() bool {
	return t.status == domain.Err
}

func (t *Task) IsPending() bool {
	return t.status == domain.Pending
}

func NewTask() domain.Task {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	return &Task{
		wg:     wg,
		data:   "",
		status: domain.Pending,
	}
}
