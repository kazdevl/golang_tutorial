package sorting_algorithm

func BubbleSort(data []int) {
	for loop_num := 0; loop_num < len(data); loop_num++ {
		for index := 0; index < len(data)-1-loop_num; index++ {
			if data[index] > data[index+1] {
				tmp := data[index]
				data[index] = data[index+1]
				data[index+1] = tmp
			}
		}
	}
}
