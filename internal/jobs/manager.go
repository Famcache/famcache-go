package jobs

import (
	"net"

	"github.com/famcache/famcache-go/domain"
)

type JobsManager struct {
	jobs     map[string]domain.Job
	socket   *net.Conn
	commands domain.Commands
}

func NewJobsManager(commands domain.Commands) domain.JobsManager {
	return &JobsManager{
		jobs:     make(map[string]domain.Job),
		commands: commands,
	}
}
