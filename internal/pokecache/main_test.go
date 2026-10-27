package pokecache

import (
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	cache.Add("testKey", []byte("some bytes"))
	data, found := cache.Get("testKey")

	if !found {
		t.Fatal("expected testKey to exist in the cache")
	}

	if string(data) != "some bytes" {
		t.Errorf("expected %q, got %q", "some bytes", string(data))
	}
}
