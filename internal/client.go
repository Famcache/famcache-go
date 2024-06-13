package internal

import (
	"net"

	"github.com/famcache/famcache-go/domain"
)

type FamcacheClient struct {
	messaging domain.Messaging
	commands  domain.Commands
	socket    net.Conn
	options   domain.Options

	taskRegistry domain.Registry
}

func NewClient(host string, port int) domain.Client {
	commands := NewCommands()

	return &FamcacheClient{
		messaging:    NewMessaging(commands),
		options:      NewOptions(host, port),
		commands:     commands,
		taskRegistry: NewRegistry(),
	}
}
