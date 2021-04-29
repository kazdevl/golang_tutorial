package sorting_algorithm

func SelectionSort(data []int) {
	for loop_num := 0; loop_num < len(data); loop_num++ {
		max_value := struct {
			value int
			index int
		}{
			value: data[loop_num],
			index: loop_num,
		}
		for index := loop_num; index < len(data); index++ {
			if max_value.value < data[index] {
				max_value.value = data[index]
				max_value.index = index
			}
		}
		tmp := data[loop_num]
		data[loop_num] = max_value.value
		data[max_value.index] = tmp
	}
}
