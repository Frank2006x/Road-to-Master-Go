package main

import (
	"fmt"
)

func main() {
	// MapWithMutex example
	fmt.Println("=== MapWithMutex ===")
	mMutex := NewMapWithMutex()
	mMutex.Set("key1", 100)
	mMutex.Set("key2", 200)
	if val, ok := mMutex.Get("key1"); ok {
		fmt.Printf("key1: %d\n", val)
	}
	mMutex.Delete("key2")
	fmt.Println("Deleted key2")

	// SyncMap example
	fmt.Println("\n=== SyncMap ===")
	mSync := NewSyncMap()
	mSync.Set("key1", 100)
	mSync.Set("key2", 200)
	if val, ok := mSync.Get("key1"); ok {
		fmt.Printf("key1: %d\n", val)
	}
	mSync.Delete("key2")
	fmt.Println("Deleted key2")

	// RawMap example
	fmt.Println("\n=== RawMap ===")
	mRaw := NewRawMap()
	mRaw.Set("key1", 100)
	mRaw.Set("key2", 200)
	if val, ok := mRaw.Get("key1"); ok {
		fmt.Printf("key1: %d\n", val)
	}
	mRaw.Delete("key2")
	fmt.Println("Deleted key2")

	fmt.Println("\nRun benchmarks with: go test -bench=. -benchmem")
}
