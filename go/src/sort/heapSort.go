package sort

// HeapSort
func heapSort(c []int) []int {
	n := len(c)

	// init heap
	for start := n / 2; start >= 0; start-- {
		maxHeap(start, n-1, c)
	}

	// sort
	for end := n - 1; end > 0; end-- {
		c[0], c[end] = c[end], c[0]
		maxHeap(0, end-1, c)
	}
	return c
}

// percolateDown
func maxHeap(start int, end int, c []int) {
	root := start
	for true {
		child := 2*root + 1
		if child > end {
			break
		}
		if child+1 < end && c[child] < c[child+1] {
			child++
		}
		if c[root] < c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
	return
}
