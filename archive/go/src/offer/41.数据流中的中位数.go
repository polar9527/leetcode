package main

type MedianFinder struct {
	minHeap []int
	maxHeap []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	min := make([]int, 0)
	max := make([]int, 0)

	return MedianFinder{min, max}
}

func (this *MedianFinder) AddNum(num int) {
	// 第一个元素，扔到小顶堆
	if len(this.minHeap) == 0 {
		this.minHeap = append(this.minHeap, num)
		return
	}

	if len(this.minHeap) < len(this.maxHeap) {
		// 添加到小顶堆
		if num < this.maxHeap[0] {
			this.minHeap = append(this.minHeap, this.maxHeap[0])
			this.minHeapUp()
			this.maxHeap[0] = num
			this.maxHeapDown()
		} else {
			this.minHeap = append(this.minHeap, num)
			this.minHeapUp()
		}
	} else {
		// 添加到大顶堆
		if num > this.minHeap[0] {
			this.maxHeap = append(this.maxHeap, this.minHeap[0])
			this.maxHeapUp()
			this.minHeap[0] = num
			this.minHeapDown()
		} else {
			this.maxHeap = append(this.maxHeap, num)
			this.maxHeapUp()
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.minHeap) == 0 || len(this.maxHeap) == 0 {
		return -1
	}
	if len(this.minHeap) == len(this.maxHeap) {
		return float64(this.minHeap[0]+this.maxHeap[0]) / 2.0
	} else if len(this.minHeap) < len(this.maxHeap) {
		return float64(this.maxHeap[0])
	} else {
		return float64(this.minHeap[0])
	}
}

func (this *MedianFinder) minHeapUp() {
	child := len(this.minHeap) - 1
	parent := (len(this.minHeap) - 1) / 2
	for child > 0 && this.minHeap[parent] > this.minHeap[child] {
		this.minHeap[child], this.minHeap[parent] = this.minHeap[parent], this.minHeap[child]
		child = parent
		parent = (parent - 1) / 2
	}
}

func (this *MedianFinder) minHeapDown() {
	parent := 0
	childL := 2*parent + 1
	childR := childL + 1

	for childL < len(this.minHeap) {
		min := childL
		if childR < len(this.minHeap) && this.minHeap[childL] > this.minHeap[childR] {
			min = childR
		}
		if this.minHeap[min] > this.minHeap[parent] {
			min = parent
		}
		if min == parent {
			break
		}
		this.minHeap[min], this.minHeap[parent] = this.minHeap[parent], this.minHeap[min]
		parent = min
		childL = 2*parent + 1
		childR = childL + 1
	}
}

func (this *MedianFinder) maxHeapUp() {
	child := len(this.maxHeap) - 1
	parent := (len(this.maxHeap) - 1) / 2
	for child > 0 && this.maxHeap[parent] < this.maxHeap[child] {
		this.maxHeap[child], this.maxHeap[parent] = this.maxHeap[parent], this.maxHeap[child]
		child = parent
		parent = (parent - 1) / 2
	}
}

func (this *MedianFinder) maxHeapDown() {
	parent := 0
	childL := 2*parent + 1
	childR := childL + 1

	for childL < len(this.maxHeap) {
		min := childL
		if childR < len(this.maxHeap) && this.maxHeap[childL] < this.maxHeap[childR] {
			min = childR
		}
		if this.maxHeap[min] < this.maxHeap[parent] {
			min = parent
		}
		if min == parent {
			break
		}
		this.maxHeap[min], this.maxHeap[parent] = this.maxHeap[parent], this.maxHeap[min]
		parent = min
		childL = 2*parent + 1
		childR = childL + 1
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	f := Constructor()
	f.FindMedian()
	f.AddNum(1)
	f.AddNum(2)
	f.FindMedian()
	f.AddNum(3)
	f.FindMedian()

}
