package leetcode

import (
	"regexp"
	"strings"
	"unsafe"
)

func replaceSpace(s string) string {
	if len(s) == 0 || len(s) > 100000 {
		return ""
	}
	// in golang, the string is actually a constant array the length of which could not be changed, and
	// the items in this array cannot be changed either.
	// so we may need to allocate a new array to fill the result。

	// space count
	spaceCounter := 0
	for _, v := range s {
		if v == ' ' {
			spaceCounter++
		}
	}
	if spaceCounter == 0 {
		return s
	}

	// create new array
	newSB := make([]byte, spaceCounter*2+len(s))

	// get the last pointer of both arrays
	sIndex := len(s) - 1
	sbIndex := len(newSB) - 1

	// iterate over the array simultaneously
	for sIndex >= 0 {
		if s[sIndex] == ' ' {
			newSB[sbIndex] = '0'
			newSB[sbIndex-1] = '2'
			newSB[sbIndex-2] = '%'
			sbIndex -= 3
			sIndex--
		} else {
			newSB[sbIndex] = s[sIndex]
			sbIndex--
			sIndex--
		}
	}
	return *(*string)(unsafe.Pointer(&newSB))
}
func replaceSpaceRegexp(s string) string {
	if len(s) == 0 || len(s) > 100000 {
		return ""
	}
	re := regexp.MustCompile(` +?`)
	return re.ReplaceAllString(s, "%20")
}

func replaceSpaceSplit(s string) string {
	if len(s) == 0 || len(s) > 100000 {
		return ""
	}
	a := strings.Split(s, " ")
	return strings.Join(a, "%20")
}
