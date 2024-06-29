package internal

import (
	"fmt"

	"github.com/famcache/famcache-go/domain"
)

type Commands struct{}

func (c *Commands) Set(id string, key string, value string, ttl *uint64) string {
	if ttl != nil {
		return fmt.Sprintf("%s SET %s %s %d\n", id, key, value, *ttl)
	}

	return fmt.Sprintf("%s SET %s %s\n", id, key, value)
}

func (c *Commands) Get(id string, key string) string {
	return fmt.Sprintf("%s GET %s\n", id, key)
}

func (c *Commands) Delete(id string, key string) string {
	return fmt.Sprintf("%s DELETE %s\n", id, key)
}

func (c *Commands) Publish(id string, topic string, data string) string {
	return fmt.Sprintf("%s PUBLISH %s %s\n", id, topic, data)
}

func (c *Commands) Subscribe(id string, topic string) string {
	return fmt.Sprintf("%s SUBSCRIBE %s\n", id, topic)
}

func (c *Commands) Unsubscribe(id string, topic string) string {
	return fmt.Sprintf("%s UNSUBSCRIBE %s\n", id, topic)
}

func (c *Commands) RegisterJob(id string, delay uint64, periodic bool) string {
	return fmt.Sprintf("%s JOB_REGISTER %d %t\n", id, delay, periodic)
}

func (c *Commands) CancelJob(id string, jobID string) string {
	return fmt.Sprintf("%s JOB_CANCEL %s\n", id, jobID)
}

func NewCommands() domain.Commands {
	return &Commands{}
}
