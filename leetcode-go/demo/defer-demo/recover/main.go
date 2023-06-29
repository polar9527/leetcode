package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer fmt.Println("defer main")
	var user = os.Getenv("USER_")

	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err:", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()
			if user == "" {
				panic("should set user env.")
			}

			// 此处不会执行
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(100)
	fmt.Println("end of main function")
}
