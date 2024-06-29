package internal

import (
	"net"

	"github.com/famcache/famcache-go/domain"
	"github.com/famcache/famcache-go/internal/jobs"
	"github.com/famcache/famcache-go/internal/messaging"
	"github.com/famcache/famcache-go/internal/task"
)

type FamcacheClient struct {
	messaging domain.Messaging
	commands  domain.Commands
	socket    net.Conn
	jobs      domain.JobsManager
	options   domain.Options

	taskRegistry domain.Registry
}

func NewClient(host string, port int) domain.Client {
	commands := NewCommands()

	return &FamcacheClient{
		messaging:    messaging.NewMessaging(commands),
		options:      NewOptions(host, port),
		taskRegistry: task.NewRegistry(),
		jobs:         jobs.NewJobsManager(commands),
		commands:     commands,
	}
}
