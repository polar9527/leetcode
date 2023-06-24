package sort

import (
	"testing"
)

func TestRadixLSD(t *testing.T) {
	arr := []int{1, 11, 1, 343, 3, 427, 663, 2075, 43, 22}
	ret := radixLSD(arr)
	t.Log(ret)
}
