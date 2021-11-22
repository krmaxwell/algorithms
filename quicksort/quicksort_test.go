package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	unsorted := []int{10, 5, 2, 3}
	sorted := []int{2, 3, 5, 10}

	got := QuickSort(unsorted)

	if !reflect.DeepEqual(got, sorted) {
		t.Errorf("got %v want %v", got, sorted)
	}

}
