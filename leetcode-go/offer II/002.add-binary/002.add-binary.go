package offer2

import (
	"strconv"
)

// @lc code=start
func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	al, bl := len(a), len(b)
	n := max(al, bl)

	for i := 0; i < n; i++ {
		if i < al {
			carry += int(a[al-i-1] - byte('0'))
		}
		if i < bl {
			carry += int(b[bl-i-1] - byte('0'))
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
