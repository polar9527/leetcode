package main

import (
	"math/rand"
)

func getLeastNumbers(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 || len(arr) < k {
		return nil
	}
	lo := 0
	hi := len(arr) - 1
	var index = -1
	index = partition(arr, lo, hi)
	for lo <= hi {
		if k == index+1 {
			break
		}
		if k < index+1 {
			hi = index - 1
			if lo <= hi {
				index = partition(arr, lo, hi)
			}
		}
		if k > index+1 {
			lo = index + 1
			if lo <= hi {
				index = partition(arr, lo, hi)
			}
		}
	}
	return arr[:k]
}

func partition(arr []int, lo, hi int) int {
	if arr == nil || len(arr) == 0 || lo < 0 || hi > len(arr)-1 || lo > hi {
		panic("wrong partiton paramenter")
	}
	rIndex := rand.Intn(hi-lo+1) + lo
	pivot := arr[rIndex]
	arr[lo], arr[rIndex] = arr[rIndex], arr[lo]
	for lo < hi {
		for lo < hi && arr[hi] >= pivot {
			hi--
		}
		arr[lo] = arr[hi]
		for lo < hi && arr[lo] <= pivot {
			lo++
		}
		arr[hi] = arr[lo]
	}
	arr[lo] = pivot
	return lo
}

func main() {
	arr := []int{0, 0, 0, 2, 0, 5}
	getLeastNumbers(arr, 0)
}
