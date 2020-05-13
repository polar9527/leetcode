package main

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	// 窗口内最大值后面更小的数要存起来
	ret := make([]int, 0)
	d := new(Dqueue)
	for i := 0; i < k; i++ {
		for d.len() != 0 && nums[d.tail()] <= nums[i] {
			d.popBack()
		}
		d.pushBack(i)
	}
	for i := k; i < len(nums); i++ {
		ret = append(ret, nums[d.head()])
		for d.len() != 0 && nums[d.tail()] <= nums[i] {
			d.popBack()
		}
		if d.len() != 0 && i-k >= d.head() {
			d.popFront()
		}
		d.pushBack(i)
	}
	ret = append(ret, nums[d.head()])
	return ret
}

type Dqueue struct {
	data []int
}

func (d *Dqueue) len() int {
	return len(d.data)
}

func (d *Dqueue) popFront() {
	if d.len() != 0 {
		d.data = d.data[1:]
	}
}
func (d *Dqueue) popBack() {
	if d.len() != 0 {
		d.data = d.data[:len(d.data)-1]
	}
}

func (d *Dqueue) pushBack(n int) {
	d.data = append(d.data, n)
}

func (d *Dqueue) tail() int {
	if len(d.data) > 0 {
		return d.data[d.len()-1]
	}
	panic("Dqueue index out of range")
}

func (d *Dqueue) head() int {
	if len(d.data) > 0 {
		return d.data[0]
	}
	panic("Dqueue index out of range")
}

// func main(){
// 	// nums := []int{1,3,1,2,0,5} 3
// 	nums := []int{1,3,-1,-3,5,3,6,7}
// 	maxSlidingWindow(nums, 3)
//
// }
