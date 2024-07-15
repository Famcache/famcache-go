package internal

import (
	"github.com/famcache/famcache-go/internal/jobs"
	"github.com/famcache/famcache-go/internal/messaging"
	"github.com/famcache/famcache-go/internal/reply"
)

func (c *FamcacheClient) Listen() error {
	for {
		buffer := make([]byte, 1024)

		c.socket.Read(buffer)

		reply_msg := string(buffer)

		if reply.IsMessagingEvent(reply_msg) {
			messaging_msg := reply.NewMessagingReply(reply_msg)

			c.messaging.(*messaging.Messaging).Trigger(messaging_msg.Topic, messaging_msg.Data)

			continue
		}

		if reply.IsCacheReply(reply_msg) {
			event := reply.NewCacheReply(reply_msg)

			task, ok := c.taskRegistry.GetById(event.ID)

			if !ok {
				continue
			}

			task.Execute(event.Data, event.Status)
		}

		if reply.IsJobReply(reply_msg) {
			job_msg := reply.NewJobReply(reply_msg)

			c.jobs.(*jobs.JobsManager).Trigger(job_msg)
		}
	}
}
