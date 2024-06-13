package reply_test

import (
	"testing"

	"github.com/famcache/famcache-go/internal/reply"
)

func TestIsCacheReply(t *testing.T) {
	reply_msg := "ce0cfc21-fb2c-46cc-8972-b177400eeb61 OK\n"

	if !reply.IsCacheReply(reply_msg) {
		t.Errorf("Expected true, got false")
	}

	reply_msg = "MESSAGE topic data"

	if reply.IsCacheReply(reply_msg) {
		t.Errorf("Expected false, got true")
	}
}
