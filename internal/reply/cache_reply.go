package reply

import (
	"strings"

	"github.com/famcache/famcache-go/domain"
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

// <uuid> STORE <status> <data?>
func NewCacheReply(reply string) CacheReply {
	parts := normalizeReply(reply)

	data := ""

	if len(parts) > 3 {
		data = parts[3]
	}

	return CacheReply{
		ID:     parts[0],
		Status: StringToStatus(parts[2]),
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

// <uuid> STORE <status> <data>
func IsCacheReply(reply string) bool {
	parts := strings.Split(strings.ReplaceAll(reply, "\n", ""), " ")

	return parts[1] == STORE_REPLY
}
