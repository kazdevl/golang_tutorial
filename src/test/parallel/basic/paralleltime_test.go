package basic_test

import (
	"testing"
	"time"
)

func Test5sSleep(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
}

func Test10sSleep(t *testing.T) {
	t.Parallel()
	time.Sleep(10 * time.Second)
}

func Test15sSleep(t *testing.T) {
	t.Parallel()
	time.Sleep(15 * time.Second)
}

func Sleep3s() {
	time.Sleep(3 * time.Second)
}

func BenchmarkSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sleep3s()
	}
}

// go test -bench . -benchmem -benchtime 30s -v -run=Benchmark*で実行すれば処理速度が向上したが
// go test -bench . -benchmem -v -run=Benchmark*で実行すると、処理速度が改善されていなかった
func BenchmarkParallel(b *testing.B) {
	b.SetParallelism(5)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Sleep3s()
		}
	})
}
