package main

import "fmt"

func main() {

	for {
		var n int
		_, err := fmt.Scan(&n)

		if err != nil {
			break
		}

		for n > 0 {
			var m int
			_, err := fmt.Scan(&m)
			if err != nil {
				break
			}
			sum := 0
			for m > 0 {
				var num int
				_, err := fmt.Scan(&num)
				if err != nil {
					break
				}
				sum += num
				m--
			}
			fmt.Println(sum)
			if n != 1 {
				fmt.Println()
			}
			n--
		}

	}

}
