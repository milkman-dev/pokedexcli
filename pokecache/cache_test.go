package pokecache

import (
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	cache := NewCache(1 * time.Second) // Cache expires after 1 second
	key := "testKey"
	val := []byte("testValue")

	// Add item to cache
	cache.Add(key, val)

	// Retrieve the item
	retrievedVal, found := cache.Get(key)
	if !found {
		t.Fatalf("Expected to find key %s, but it was not found", key)
	}

	// Check the value
	if string(retrievedVal) != string(val) {
		t.Fatalf("Expected value %s, but got %s", val, retrievedVal)
	}
}

func TestGetNonExistentKey(t *testing.T) {
	cache := NewCache(1 * time.Second)
	key := "nonExistentKey"

	// Try to retrieve a non-existent key
	_, found := cache.Get(key)
	if found {
		t.Fatalf("Expected key %s to not exist, but it was found", key)
	}
}
