package internal

import "github.com/famcache/famcache-go/internal/reply"

func (c *FamcacheClient) Listen() error {
	for {
		buffer := make([]byte, 1024)

		c.socket.Read(buffer)

		reply_msg := string(buffer)

		if reply.IsMessagingEvent(reply_msg) {
			messaging_msg := reply.NewMessagingReply(reply_msg)

			c.messaging.(*Messaging).trigger(messaging_msg.Topic, messaging_msg.Data)

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
	}
}
