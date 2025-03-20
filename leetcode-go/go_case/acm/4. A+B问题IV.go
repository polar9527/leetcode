package main

import "fmt"

func main() {

	var n, a int

	for {
		_, err := fmt.Scan(&n)
		if err != nil || n == 0 {
			break
		}
		sum := 0
		for n > 0 {
			_, err := fmt.Scan(&a)
			if err != nil {
				break
			}
			sum += a
			n--
		}
		fmt.Println(sum)
	}
}
