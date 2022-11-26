package pool_test

import (
	"testing"

	"github.com/kazdevl/golang_tutorial/sync/pool"
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
