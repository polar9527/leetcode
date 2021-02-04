package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	testcase := [][]int{
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		{1},
		{},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		{90, 80, 60, 70, 40, 30, 20, 10},
	}

	t.Log("partitionLazySwitch")
	for i, c := range testcase {
		quickSort(c, partitionLazySwitch)
		t.Log("testcase", i, ":")
		t.Log(c)
	}
	t.Log("partitionAggressiveSwitch")
	for i, c := range testcase {
		quickSort(c, partitionAggressiveSwitch)
		t.Log("testcase", i, ":")
		t.Log(c)
	}

}
