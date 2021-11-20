package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	unsorted := []int{20, 55, 49, 74, 68, 37, 93, 94, 60, 63}
	sorted := []int{20, 37, 49, 55, 60, 63, 68, 74, 93, 94}

	got := SelectionSort(unsorted)
	if !reflect.DeepEqual(got, sorted) {
		t.Errorf("got %v want %v", got, sorted)
	}
}
