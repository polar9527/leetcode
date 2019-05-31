/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

package main

func myAtoi(str string) int {

	const int32Max = int32(^uint32(0) >> 1)
	const int32Min = ^int32Max

	var ans int64 = 0
	var sign int64 = 1
	start := false

	for _, c := range str {
		if c <= '9' && c >= '0' {
			if !start {
				start = true
			}
			ans = ans*10 + int64(c) - int64('0')
			if ans*sign > int64(int32Max) {
				return int(int32Max)
			} else if ans*sign < int64(int32Min) {
				return int(int32Min)
			}
		} else if !start && c == '+' {
			start = true
		} else if !start && c == '-' {
			start = true
			sign = -1
		} else if c == ' ' {
			if !start {
				continue
			} else {
				break
			}

		} else {
			break
		}
	}
	return int(ans * sign)
}

// func main() {
// 	s := "12356987546"
// 	ret := myAtoi(s)
// 	println(ret)
// }
