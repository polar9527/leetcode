package main

import (
	"strconv"
)

func translateNum(num int) int {

	number := strconv.Itoa(num)

	l := len(number)
	switch l {
	case 0:
		return 0
	case 1:
		return 1
	default:
	}
	counts := make([]int, l)
	counts[l-1] = 1
	tmp, _ := strconv.Atoi(number[l-2:])
	counts[l-2] = counts[l-1]
	if tmp >= 10 && tmp <= 25 {
		counts[l-2] += 1
	}

	for i := l - 3; i >= 0; i-- {
		counts[i] = counts[i+1]
		tmp, _ := strconv.Atoi(number[i : i+2])
		if tmp >= 10 && tmp <= 25 {
			counts[i] += counts[i+2]
		}
	}
	return counts[0]
}
