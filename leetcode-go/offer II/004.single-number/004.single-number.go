package offer2

// 状态转移，卡诺图
func singleNumber(nums []int) int {

	one, two := 0, 0
	for _, n := range nums {
		one = ^two & (n ^ one)
		// 下面的one是已经转换过状态的的one了
		// 即，不是初始输入的one
		two = ^one & (n ^ two)
	}
	return one
}
