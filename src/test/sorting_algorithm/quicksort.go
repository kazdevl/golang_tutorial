package sorting_algorithm

func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	pivot := data[len(data)/2]
	left := 0
	right := len(data) - 1
	for {
		for pivot < data[left] {
			left++
		}

		for pivot > data[right] {
			right--
		}

		if left >= right {
			break
		}

		tmp := data[left]
		data[left] = data[right]
		data[right] = tmp

		if data[right] == pivot {
			left++
		} else if data[left] == pivot {
			right--
		}
	}

	QuickSort(data[:left])
	QuickSort(data[right+1:])
}
