package leetcode

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

import (
	"strconv"
)

func countAndSay(n int) string {
	curStr := []byte("1")
	if n == 1 {
		return string(curStr)
	}
	length := 1
	tempStr := []byte{}

	for i := 2; i <= n; i++ {

		for j := 0; j < len(curStr)-1; j++ {
			if curStr[j] == curStr[j+1] {
				length++
			} else {
				tempStr = append(tempStr, []byte(strconv.Itoa(length))...)
				tempStr = append(tempStr, curStr[j])
				length = 1
			}
		}
		tempStr = append(tempStr, []byte(strconv.Itoa(length))...)
		tempStr = append(tempStr, curStr[len(curStr)-1])
		length = 1
		curStr = tempStr
		tempStr = []byte{}
	}
	return string(curStr)
}

// func main() {
// 	fmt.Println(countAndSay(5))
// 	// i := strconv.Itoa(12345)

// 	// i := "1234"[2]
// 	// fmt.Println(reflect.TypeOf(i))
// }
