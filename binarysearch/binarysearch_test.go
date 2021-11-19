package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}

	cases := []struct {
		Key   int
		Want  int
		Error error
	}{
		{6, 5, nil},
		{1, 0, nil},
		{99, 0, ErrNotFound},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("found %d at %d", test.Key, test.Want), func(t *testing.T) {
			got, err := BinarySearch(data, test.Key)

			if got != test.Want {
				t.Errorf("got %d want %d", got, test.Want)
			}

			if err != test.Error {
				t.Errorf("got error %q want %q", err, test.Error)
			}
		})
	}

	t.Run("unsorted data", func(t *testing.T) {
		unsortedData := []int{1, 3, 2, 5, 4}
		_, err := BinarySearch(unsortedData, 3)
		wantErr := ErrNotSorted

		if err != wantErr {
			t.Errorf("got error %q want %q", err, wantErr)
		}
	})
}
