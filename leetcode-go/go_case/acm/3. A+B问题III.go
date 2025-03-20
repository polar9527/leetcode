package main

import "fmt"

func main() {
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)

		if err != nil {
			break
		}
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}

}
