package internal_test

import (
	"testing"

	"github.com/famcache/famcache-go/internal"
)

func TestRegistrySetAndGet(t *testing.T) {
	registry := internal.NewRegistry()

	registry.Set("1", internal.NewTask())

	task, ok := registry.GetById("1")

	if !ok {
		t.Errorf("Expected task to be found")
	}

	if task == nil {
		t.Errorf("Expected task to be not nil")
	}
}

func TestRegistryFree(t *testing.T) {
	registry := internal.NewRegistry()

	registry.Set("1", internal.NewTask())

	registry.Free("1")

	_, ok := registry.GetById("1")

	if ok {
		t.Errorf("Expected task to be not found")
	}
}
