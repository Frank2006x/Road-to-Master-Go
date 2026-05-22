package main

import (
	"fmt"
	"testing"
)

// Benchmark MapWithMutex
func BenchmarkMapWithMutex_Set(b *testing.B) {
	m := NewMapWithMutex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(fmt.Sprintf("key_%d", i%1000), i)
	}
}

func BenchmarkMapWithMutex_Get(b *testing.B) {
	m := NewMapWithMutex()
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key_%d", i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(fmt.Sprintf("key_%d", i%1000))
	}
}

func BenchmarkMapWithMutex_SetGet(b *testing.B) {
	m := NewMapWithMutex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(fmt.Sprintf("key_%d", i%1000), i)
		m.Get(fmt.Sprintf("key_%d", i%1000))
	}
}

// Benchmark SyncMap
func BenchmarkSyncMap_Set(b *testing.B) {
	m := NewSyncMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(fmt.Sprintf("key_%d", i%1000), i)
	}
}

func BenchmarkSyncMap_Get(b *testing.B) {
	m := NewSyncMap()
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key_%d", i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(fmt.Sprintf("key_%d", i%1000))
	}
}

func BenchmarkSyncMap_SetGet(b *testing.B) {
	m := NewSyncMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(fmt.Sprintf("key_%d", i%1000), i)
		m.Get(fmt.Sprintf("key_%d", i%1000))
	}
}

// Benchmark RawMap
func BenchmarkRawMap_Set(b *testing.B) {
	m := NewRawMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(fmt.Sprintf("key_%d", i%1000), i)
	}
}

func BenchmarkRawMap_Get(b *testing.B) {
	m := NewRawMap()
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key_%d", i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(fmt.Sprintf("key_%d", i%1000))
	}
}

func BenchmarkRawMap_SetGet(b *testing.B) {
	m := NewRawMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(fmt.Sprintf("key_%d", i%1000), i)
		m.Get(fmt.Sprintf("key_%d", i%1000))
	}
}

// Parallel benchmarks for concurrent scenarios
func BenchmarkMapWithMutex_Concurrent(b *testing.B) {
	m := NewMapWithMutex()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%3 == 0 {
				m.Set(fmt.Sprintf("key_%d", i%1000), i)
			} else if i%3 == 1 {
				m.Get(fmt.Sprintf("key_%d", i%1000))
			} else {
				m.Delete(fmt.Sprintf("key_%d", i%1000))
			}
			i++
		}
	})
}

func BenchmarkSyncMap_Concurrent(b *testing.B) {
	m := NewSyncMap()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%3 == 0 {
				m.Set(fmt.Sprintf("key_%d", i%1000), i)
			} else if i%3 == 1 {
				m.Get(fmt.Sprintf("key_%d", i%1000))
			} else {
				m.Delete(fmt.Sprintf("key_%d", i%1000))
			}
			i++
		}
	})
}

func BenchmarkRawMap_Concurrent(b *testing.B) {
	m := NewRawMap()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%3 == 0 {
				m.Set(fmt.Sprintf("key_%d", i%1000), i)
			} else if i%3 == 1 {
				m.Get(fmt.Sprintf("key_%d", i%1000))
			} else {
				m.Delete(fmt.Sprintf("key_%d", i%1000))
			}
			i++
		}
	})
}
