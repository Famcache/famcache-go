package domain

type Registry interface {
	GetById(id string) (Task, bool)
	Set(id string, task Task)
	Free(id string)
}
