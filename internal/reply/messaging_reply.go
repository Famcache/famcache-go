package reply

// ─────────────────────────────────────────────────────────────────────────────

const MessagingReplyPrefix = "MESSAGE"

// ─────────────────────────────────────────────────────────────────────────────

type MessagingReply struct {
	Topic string
	Data  string
}

func NewMessagingReply(reply string) MessagingReply {
	parts := normalizeReply(reply)

	return MessagingReply{
		Topic: parts[1],
		Data:  parts[2],
	}
}

// ─────────────────────────────────────────────────────────────────────────────

func IsMessagingEvent(reply string) bool {
	return len(reply) >= len(MessagingReplyPrefix) && reply[:len(MessagingReplyPrefix)] == MessagingReplyPrefix
}
