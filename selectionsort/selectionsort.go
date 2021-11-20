package selectionsort

func SelectionSort(data []int) []int {
	var sorted []int
	var min_index int

	data_length := len(data)
	for i := 0; i < data_length; i++ {
		min_index = findSmallest(data)
		sorted = append(sorted, data[min_index])
		data = removeSliceElement(data, min_index)
	}

	return sorted
}

func removeSliceElement(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func findSmallest(data []int) int {
	minimum := data[0]
	min_index := 0
	for i := 1; i < len(data); i++ {
		if data[i] < minimum {
			minimum = data[i]
			min_index = i
		}
	}
	return min_index
}
