package binarysearch

const (
	ErrNotFound  = SearchError("not found")
	ErrNotSorted = SearchError("data not sorted")
)

type SearchError string

func (e SearchError) Error() string {
	return string(e)
}

func BinarySearch(data []int, search int) (int, error) {
	low := 0
	high := len(data) - 1
	var mid int
	var check int

	for low <= high {
		if data[low] > data[high] {
			return 0, ErrNotSorted
		}
		mid = (low + high) / 2
		check = data[mid]
		if check == search {
			return mid, nil
		}
		if check > search {
			high = mid - 1
		}
		if check < search {
			low = mid + 1
		}
	}
	return 0, ErrNotFound // TODO: return an error
}
