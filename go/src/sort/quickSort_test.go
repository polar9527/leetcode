package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	var a = []int{9, 8, 6, 7, 4, 3, 2, 1}
	quickSort(a, partitionA)
	t.Log(a)

	var l = []int{90, 80, 60, 70, 40, 30, 20, 10}
	quickSort(l, partitionL)
	t.Log(l)

}
