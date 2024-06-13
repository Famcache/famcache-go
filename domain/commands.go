package domain

type Commands interface {
	Set(id string, key string, value string, ttl *uint64) string
	Get(id string, key string) string
	Delete(id string, key string) string

	Publish(id string, topic string, data string) string
	Subscribe(id string, topic string) string
	Unsubscribe(id string, topic string) string
}
