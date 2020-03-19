package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "fasdf  sd af"
	ret := replaceSpace(s)
	fmt.Println(ret)
}

func replaceSpace(s string) string {
	re := regexp.MustCompile(` +?`)
	return re.ReplaceAllString(s, "%20")
}
