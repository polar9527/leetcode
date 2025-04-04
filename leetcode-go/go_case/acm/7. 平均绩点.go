package acm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		list := strings.Split(string(s), " ")

		sum := 0.0
		flag := true
		for i := 0; i < len(list) && flag; i++ {
			switch list[i] {
			case "A":
				sum += 4.0
			case "B":
				sum += 3.0
			case "C":
				sum += 2.0
			case "D":
				sum += 1.0
			case "F":
				continue
			default:
				flag = false
			}
		}
		if flag {
			fmt.Printf("%.2f\n", sum)
		} else {
			fmt.Println("Unknown")
		}
	}
}
