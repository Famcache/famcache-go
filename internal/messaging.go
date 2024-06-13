package internal

import (
	"net"

	"github.com/famcache/famcache-go/domain"
	"github.com/google/uuid"
)

type Messaging struct {
	subscriptions map[string]chan string
	socket        *net.Conn
	commands      domain.Commands
}

func (m *Messaging) Publish(topic string, message string) error {
	uuid := uuid.New().String()

	command := m.commands.Publish(uuid, topic, message)

	_, err := (*m.socket).Write([]byte(command))

	return err
}

func (m *Messaging) Subscribe(topic string) (<-chan string, error) {
	uuid := uuid.New().String()

	command := m.commands.Subscribe(uuid, topic)

	_, err := (*m.socket).Write([]byte(command))

	if err != nil {
		return nil, err
	}

	m.subscriptions[topic] = make(chan string)

	return m.subscriptions[topic], nil
}

func (m *Messaging) trigger(topic string, data string) {
	if ch, ok := m.subscriptions[topic]; ok {
		ch <- data
	}
}

func (m *Messaging) Unsubscribe(topic string) error {
	uuid := uuid.New().String()

	command := m.commands.Unsubscribe(uuid, topic)

	_, err := (*m.socket).Write([]byte(command))

	if err != nil {
		return err
	}

	delete(m.subscriptions, topic)

	return nil
}

func NewMessaging(commands domain.Commands) domain.Messaging {
	return &Messaging{
		subscriptions: make(map[string]chan string),
		commands:      commands,
	}
}
