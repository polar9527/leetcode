package main

import (
	"sort"
	"strconv"
	"strings"
)

type IntSlice []int

func (s IntSlice) Len() int      { return len(s) }
func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool {
	iString := strconv.Itoa(s[i])
	jString := strconv.Itoa(s[j])
	// iString + jString
	// 可能溢出
	// ij, _ := strconv.Atoi(iString+jString)
	// ji, _ := strconv.Atoi(jString+iString)
	// 	if ij < ji {
	//		return true
	//	} else {
	//		return false
	//	}
	ijString := iString + jString
	jiString := jString + iString

	for k := 0; k < len(ijString); k++ {
		if ijString[k] == jiString[k] {
			continue
		} else {
			if ijString[k] < jiString[k] {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func minNumber(nums []int) string {
	sort.Sort(IntSlice(nums))
	numsString := make([]string, 0)
	for _, n := range nums {
		numsString = append(numsString, strconv.Itoa(n))
	}
	return strings.Join(numsString, "")
}

// func main() {
// 	nums := []int{3, 30}
// 	minNumber(nums)
// }
