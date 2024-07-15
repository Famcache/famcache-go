package task

import "github.com/famcache/famcache-go/domain"

type Registry struct {
	tasks map[string]domain.Task
}

func (r *Registry) GetById(id string) (domain.Task, bool) {
	task, ok := r.tasks[id]

	return task, ok
}

func (r *Registry) Set(id string, task domain.Task) {
	r.tasks[id] = task
}

func (r *Registry) Free(id string) {
	delete(r.tasks, id)
}

func NewRegistry() domain.Registry {
	return &Registry{
		tasks: make(map[string]domain.Task),
	}
}
