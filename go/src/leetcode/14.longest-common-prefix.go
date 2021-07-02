/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]
	for i := 1; i < len(strs); i++ {
		second := strs[i]
		first = getPrefix(first, second)
	}
	return first
}

func getPrefix(f, s string) string {
	ret := ""
	l := 0
	if len(f) < len(s) {
		l = len(f)
	} else {
		l = len(s)
	}
	index := -1
	for i := 0; i < l; i++ {
		if f[i] == s[i] {
			index = i
		} else {
			break
		}
	}
	if index < 0 {
		ret = ""
	} else {
		ret = f[:index+1]
	}

	return ret

}

// import (
//     "strings"
// )
// func longestCommonPrefix(strs []string) string {
//     if len(strs) == 0 {
//         return ""
//     }
//     var b strings.Builder
//     for index :=0; index < len(strs[0]); index++ {
//         for i := 1; i < len(strs); i++ {
//             if index >= len(strs[i-1]) || index >= len(strs[i]) || strs[i-1][index] != strs[i][index] {
//                 return b.String()
//             }
//         }
//         b.WriteString(string(strs[0][index]))
//     }
//     return b.String()
// }

// func main() {
// 	// strs := []string{"flower", "flow", "flight"}
// 	strs := []string{"c", "c"}
// 	ret := longestCommonPrefix(strs)
// 	fmt.Println(ret)
// }
