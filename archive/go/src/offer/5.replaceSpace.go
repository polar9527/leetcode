package space

import (
	"regexp"
)

func replaceSpaceCheat(s string) string {
	re := regexp.MustCompile(` +?`)
	return re.ReplaceAllString(s, "%20")
}

package bintree

import (
	"fmt"
	"regexp"
)

func replaceSpaceCheat(s string) string {
	re := regexp.MustCompile(` +?`)
	return re.ReplaceAllString(s, "%20")
}

func replaceSpace(s string) string {
	originLen := len(s)
	if originLen == 0 {
		return s
	}

	spaceCounter := 0
	for _, r := range s {
		if r == ' ' {
			spaceCounter++
		}
	}
	// newLen := originLen + spaceCounter*3

	p1 := originLen - 1
	// len("%20") - len(" ") = 2
	for i := 0; i < spaceCounter*2; i++ {
		s = s + "#"
	}
	p2 := len(s) - 1
	// b := []byte(s)
	b := *(*[]byte)(unsafe.Pointer(&s))
	bPtr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sPtr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(bPtr)
	fmt.Println(sPtr)
	for p1 >= 0 {
		logs := string(b)
		fmt.Sprintln(logs)
		if s[p1] == ' ' {
			p1--
			b[p2] = '0'
			b[p2-1] = '2'
			b[p2-2] = '%'
			p2 -= 3
		} else {
			b[p2] = b[p1]
			p1--
			p2--
		}
	}
	return string(b)
}