package sort

// HeapSort
func heapSort(c []int) []int {
	n := len(c)

	// init heap
	for start := n / 2; start >= 0; start-- {
		// 每一个非叶子节点做一次下滤，从最低的非叶子节点开始，直到最顶层节点
		maxHeap(start, n-1, c)
	}

	// sort
	for end := n - 1; end > 0; end-- {
		// 取出最大元素，缩小堆，然后重新下滤一次
		c[0], c[end] = c[end], c[0]
		maxHeap(0, end-1, c)
	}
	return c
}

// percolateDown
func maxHeap(start int, end int, c []int) {
	root := start
	for root <= end {
		child := 2*root + 1
		if child+1 <= end && c[child] < c[child+1] {
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
