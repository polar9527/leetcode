package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("enter defer 1")
		defer func() {
			fmt.Println("enter defer 2")
			// recover()
			// https://stackoverflow.com/questions/72651899/why-golang-can-not-recover-from-a-panic-in-a-function-called-by-the-defer-functi
			if e := recover(); e != nil {
				fmt.Println("recover")
			}
			fmt.Println("exit defer 2")
		}()
		fmt.Println("exit defer 1")
	}()
	panic(404)
}
