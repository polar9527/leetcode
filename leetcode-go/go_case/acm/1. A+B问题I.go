package acm

import "fmt"

func main() {

	var a, b int
	for {

		_, err := fmt.Scanf("%d %d", &a, &b)

		if err != nil {
			break
		}
		fmt.Println(a + b)
	}
}
