package quicksort

func QuickSort(data []int) []int {
	if len(data) < 2 {
		return data
	} else {
		// TODO: choose a better pivot point (median of 3 random elements??)
		pivot_point := len(data) / 2
		pivot := data[pivot_point]
		var lesser []int
		var greater []int
		for i := 0; i < len(data); i++ {
			if i == pivot_point {
				continue
			} else if data[i] <= pivot {
				lesser = append(lesser, data[i]) // TODO: is this efficient?
			} else {
				greater = append(greater, data[i])
			}
		}
		result := QuickSort(lesser)
		result = append(result, pivot) // TODO: can we combine the two halves more efficiently?
		result = append(result, QuickSort(greater)...)
		return result
	}
}
