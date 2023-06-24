package sort

import (
	"testing"
)

func TestShellSort(t *testing.T) {

	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	shellSort(arr)
	t.Log(arr)
}
