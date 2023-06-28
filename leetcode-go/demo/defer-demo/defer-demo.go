package demo

import "fmt"

func DeferDemo(flag bool) {
	fmt.Println(flag)
	re := f()
	fmt.Println(re)
}

func f() (r int) {
	defer func() {
		r = r + 5
	}()
	return 1
}
