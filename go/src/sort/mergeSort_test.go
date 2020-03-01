package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	t.Log(mergeSort(arr))
}
