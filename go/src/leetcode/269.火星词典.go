/*
* 现有一种使用字母的全新语言，这门语言的字母顺序与英语顺序不同。
*
* 假设，您并不知道其中字母之间的先后顺序。但是，会收到词典中获得一个 不为空的 单词列表。因为是从词典中获得的，所以该单词列表内的单词已经 按这门新语言的字母顺序进行了排序。
*
* 您需要根据这个输入的列表，还原出此语言中已知的字母顺序。
*
*
*
* 示例 1：
*
* 输入:
* [
* "wrt",
* "wrf",
* "er",
* "ett",
* "rftt"
* ]
* 输出: "wertf"
* 示例 2：
*
* 输入:
* [
* "z",
* "x"
* ]
* 输出: "zx"
* 示例 3：
*
* 输入:
* [
* "z",
* "x",
* "z"
* ]
* 输出: ""
* 解释: 此顺序是非法的，因此返回 ""。
*
*
* 提示：
*
* 你可以默认输入的全部都是小写字母
* 若给定的顺序是不合法的，则返回空字符串即可
* 若存在多种可能的合法字母顺序，请返回其中任意一种顺序即可
*
* 来源：力扣（LeetCode）
* 链接：https://leetcode-cn.com/problems/alien-dictionary
* 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

package main

import (
	"fmt"
)

func alienOrder(words []string) string {
	var ansByte []byte

	// dict
	edgesOfNodes := make(map[byte]map[byte]bool)
	ingress := make(map[byte]int)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if _, ok := edgesOfNodes[word[i]]; !ok {
				edgesOfNodes[word[i]] = make(map[byte]bool)
				ingress[word[i]] = 0
			}
		}
	}

	// edges := make(map[byte][]byte)
	n := len(words)
	for i := 0; i < n-1; i++ {
		index := 0
		for ; index < len(words[i]) && index < len(words[i+1]); index++ {
			if words[i][index] != words[i+1][index] {
				if _, ok := edgesOfNodes[words[i][index]][words[i+1][index]]; !ok {
					edgesOfNodes[words[i][index]][words[i+1][index]] = true
					ingress[words[i+1][index]] += 1
				}
				break
			}
		}
		if index == len(words[i+1]) && len(words[i]) > len(words[i+1]) {
			return ""
		}
	}

	queue := make([]byte, 0)
	for k, v := range ingress {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		ansByte = append(ansByte, cur)
		for neighbor := range edgesOfNodes[cur] {
			ingress[neighbor]--
			if ingress[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if len(ansByte) != len(edgesOfNodes) {
		return ""
	}
	return string(ansByte)
}

func main() {
	// ss := []string{"wrt","wrf","er","ett","rftt","te"}
	ss := []string{"ozvcdpgfq", "mridvkklqj", "dpwecbwor", "xxtistijm", "xxuon", "tudbazpggu", "hnuumktbjy", "bogbcoi"}
	fmt.Println(alienOrder(ss))
}
