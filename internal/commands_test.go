package internal_test

import (
	"testing"

	"github.com/famcache/famcache-go/internal"
)

func TestGetCommand(t *testing.T) {
	commands := internal.NewCommands()

	expected := "1 GET key\n"

	if got := commands.Get("1", "key"); got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestSetCommandWithoutTTL(t *testing.T) {
	commands := internal.NewCommands()

	expected := "1 SET key value\n"

	if got := commands.Set("1", "key", "value", nil); got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestSetCommandWithTTL(t *testing.T) {
	commands := internal.NewCommands()

	expected := "1 SET key value 100\n"

	ttl := uint64(100)

	if got := commands.Set("1", "key", "value", &ttl); got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestDeleteCommand(t *testing.T) {
	commands := internal.NewCommands()

	expected := "1 DELETE key\n"

	if got := commands.Delete("1", "key"); got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestPublishCommand(t *testing.T) {
	commands := internal.NewCommands()

	expected := "1 PUBLISH topic data\n"

	if got := commands.Publish("1", "topic", "data"); got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestSubscribeCommand(t *testing.T) {
	commands := internal.NewCommands()

	expected := "1 SUBSCRIBE topic\n"

	if got := commands.Subscribe("1", "topic"); got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
