package sorting_algorithm_test

import (
	sa "app/test/sorting_algorithm"
	"math/rand"
	"testing"
	"time"
)

func makeRandomOrderList(length int) []int {
	data := make([]int, length)
	for i := 0; i < length; i++ {
		data[i] = i
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		target_index := rand.Intn(length)
		tmp := data[i]
		data[i] = data[target_index]
		data[target_index] = tmp
	}
	return data
}

func checkAssendingOrder(data []int) bool {
	for i := 0; i < len(data)-1; i++ {
		if data[i+1] > data[i] {
			return false
		}
	}
	return true
}

func BenchmarkSorts(b *testing.B) {
	var original_data []int = makeRandomOrderList(3000)
	sorts := []struct {
		name     string
		function func([]int)
	}{
		{
			name:     "bubble sort",
			function: sa.BubbleSort,
		},
		{
			name:     "selection sort",
			function: sa.SelectionSort,
		},
		{
			name:     "quick sort",
			function: sa.QuickSort,
		},
	}
	for _, sort := range sorts {
		data := make([]int, len(original_data))
		copy(data, original_data)

		b.Run(sort.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sort.function(data)
			}
			b.Cleanup(func() {
				b.Logf("%v is finished. result is %v\n", sort.name, checkAssendingOrder(data))
			})
		})
	}
}
