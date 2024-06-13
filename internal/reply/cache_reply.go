package reply

import (
	"strings"

	"github.com/famcache/famcache-go/domain"
	"github.com/google/uuid"
)

// ─────────────────────────────────────────────────────────────────────────────

const (
	StatusOk  = "OK"
	StatusErr = "ERROR"
)

type CacheReply struct {
	ID     string
	Status domain.Status
	Data   string
}

func NewCacheReply(reply string) CacheReply {
	parts := normalizeReply(reply)

	data := ""

	if len(parts) > 2 {
		data = parts[2]
	}

	return CacheReply{
		ID:     parts[0],
		Status: StringToStatus(parts[1]),
		Data:   data,
	}
}

// ─────────────────────────────────────────────────────────────────────────────

func StringToStatus(status string) domain.Status {
	switch status {
	case "OK":
		return domain.OK
	case "ERROR":
		return domain.Err
	default:
		return domain.Err
	}
}

func IsCacheReply(reply string) bool {
	parts := strings.Split(strings.ReplaceAll(reply, "\n", ""), " ")

	if len(parts) < 2 {
		return false
	}

	_, err := uuid.Parse(parts[0])

	return err == nil
}
