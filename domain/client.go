package domain

type Client interface {
	Connect() error
	Close() error

	Set(key string, value string, ttl *uint64) error
	Get(key string) (string, error)
	Delete(key string) error

	Messaging() Messaging
}
