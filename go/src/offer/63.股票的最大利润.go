package main

func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	min := prices[0]
	p := 0
	for i := 1; i < len(prices); i++ {
		p = max(p, prices[i]-min)
		if prices[i] < min {
			min = prices[i]
		}
	}
	return p
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	maxProfit(nums)
}
