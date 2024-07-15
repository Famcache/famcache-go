package messaging

import (
	"net"

	"github.com/famcache/famcache-go/domain"
)

type Messaging struct {
	subscriptions map[string]chan string
	socket        *net.Conn
	commands      domain.Commands
}

func NewMessaging(commands domain.Commands) domain.Messaging {
	return &Messaging{
		subscriptions: make(map[string]chan string),
		commands:      commands,
	}
}
