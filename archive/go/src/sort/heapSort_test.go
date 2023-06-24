package sort

import (
	"testing"
)

// test
func TestHeapSort(t *testing.T) {
	var b = []int{9, 8, 6, 7, 4, 3, 2, 1}
	heapSort(b)
	t.Log(b)
}
