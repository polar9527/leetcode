package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	selectionSort(arr)
	t.Log(arr)
}
