package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	insertionSort(arr)
	t.Log(arr)
}
