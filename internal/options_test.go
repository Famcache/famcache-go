package internal_test

import (
	"testing"

	"github.com/famcache/famcache-go/internal"
)

func TestNewOptions(t *testing.T) {
	options := internal.NewOptions("localhost", 8080)

	if options.GetHost() != "localhost" {
		t.Errorf("Expected host to be localhost, got %s", options.GetHost())
	}

	if options.GetPort() != 8080 {
		t.Errorf("Expected port to be 8080, got %d", options.GetPort())
	}
}
