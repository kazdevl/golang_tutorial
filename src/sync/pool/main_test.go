package pool_test

import (
	"app/sync/pool"
	"testing"
)

func Benchmark_CreateListEach(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.CreateListEach()
	}
}

func Benchmark_CreateListOnce(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.CreateListOnce()
	}
}

func Benchmark_UseSyncPool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.UseSyncPool()
	}
}
