package domain

type Messaging interface {
	Publish(topic string, message string) error
	Subscribe(topic string) (<-chan string, error)
	Unsubscribe(topic string) error
}
