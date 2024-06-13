package reply_test

import (
	"testing"

	"github.com/famcache/famcache-go/internal/reply"
)

func TestIsMessagingReply(t *testing.T) {
	reply_msg := "MESSAGE topic data"

	if !reply.IsMessagingEvent(reply_msg) {
		t.Errorf("Expected true, got false")
	}

	reply_msg = "uuid-uuid-uuid OK data"

	if reply.IsMessagingEvent(reply_msg) {
		t.Errorf("Expected false, got true")
	}

}
