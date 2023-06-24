package main

import (
	"strconv"
)

func translateNum(num int) int {

	number := strconv.Itoa(num)

	return core([]byte(number))

}

func core(num []byte) int {

	l := len(num)
	if l == 0 {
		return -1
	}

	countArr := make([]int, l)

	for count, i := 0, l-1; i >= 0; i-- {
		if i == l-1 {
			count = 1
		} else {
			count = countArr[i+1]
		}

		if i < l-1 {
			var convert []byte
			convert = append(convert, num[i:i+2]...)
			number, _ := strconv.Atoi(string(convert))
			if number >= 10 && number <= 25 {
				if i < l-2 {
					count += countArr[i+2]
				} else {
					count += 1
				}
			}
		}

		countArr[i] = count
	}

	// countArr[l-1] = 1
	// if l >= 2 {
	// 	var convert []byte
	// 	convert = append(convert, num[l-2:]...)
	// 	number, _ := strconv.Atoi(string(convert))
	// 	if number >= 10 && number <= 25 {
	// 		countArr[l-2] = 2
	// 	} else {
	// 		countArr[l-2] = 1
	// 	}
	// 	if l >= 3 {
	// 		for i := l - 3; i >= 0; i--{
	// 			var convert []byte
	// 			convert = append(convert, num[i:i+2]...)
	// 			number, _ := strconv.Atoi(string(convert))
	// 			if number >= 10 && number <= 25 {
	// 				countArr[i] = countArr[i+1] + countArr[i+2]
	// 			} else {
	// 				countArr[i] = countArr[i+1]
	// 			}
	// 		}
	// 	}
	//
	// }

	return countArr[0]
}

func main() {
	translateNum(25)
}
