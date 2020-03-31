package main

import (
	"fmt"
)

func printNumbers(n int) []int {
	numbers := make([]int, 0)
	if n <= 0 {
		return numbers
	}
	numberBytes := make([]byte, n)

	for i := range numberBytes {
		numberBytes[i] = '0'
	}

	for !increment(numberBytes) {
		number := getNumber(numberBytes)
		numbers = append(numbers, number)
	}

	return numbers
}

func increment(numberBytes []byte) bool {
	length := len(numberBytes)
	isOverFlow := false
	carry := 0
	for index := length - 1; index >= 0; index-- {

		sum := int(numberBytes[index]-'0') + carry
		if index == length-1 {
			sum++
		}

		if sum >= 10 {
			if index == 0 {
				isOverFlow = true
			} else {
				carry = 1
				sum -= 10
				numberBytes[index] = byte(sum) + '0'
			}
		} else {
			numberBytes[index] = byte(sum) + '0'
			break
		}

	}

	return isOverFlow
}

func getNumber(numberBytes []byte) int {
	length := len(numberBytes)
	beginning := 0
	for i := 0; i < length; i++ {
		if numberBytes[i] != '0' {
			beginning = i
			break
		}
	}

	base := 1
	sum := 0
	for i := length - 1; i >= beginning; i-- {
		sum += int(numberBytes[i]-'0') * base
		base *= 10
	}

	return sum
}

func main() {
	arr := printNumbers(2)
	fmt.Println(arr)

}
