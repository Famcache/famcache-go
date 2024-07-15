package famcache

import (
	"github.com/famcache/famcache-go/domain"
	"github.com/famcache/famcache-go/internal"
	"github.com/famcache/famcache-go/internal/jobs"
	"github.com/google/uuid"
)

func New(host string, port int) domain.Client {
	return internal.NewClient(host, port)
}

type JobConfig struct {
	Name   string
	Delay  uint64
	Repeat bool
}

func NewJob(config JobConfig) domain.Job {
	uuid := uuid.New().String()
	return jobs.NewJob(uuid, config.Name, config.Delay, config.Repeat)
}
