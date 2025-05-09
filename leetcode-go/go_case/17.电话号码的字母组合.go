package go_case

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (59.29%)
 * Likes:    2782
 * Dislikes: 0
 * Total Accepted:    840.5K
 * Total Submissions: 1.4M
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：digits = "23"
 * 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：digits = ""
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：digits = "2"
 * 输出：["a","b","c"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] 是范围 ['2', '9'] 的一个数字。
 *
 *
 */

// @lc code=start
// func letterCombinations(digits string) []string {
// 	if len(digits) == 0 {
// 		return []string{}
// 	}

// 	var s []string
// 	var path []byte
// 	letterMap := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

// 	var backtracing func(int)

// 	backtracing = func(index int) {
// 		if len(path) == len(digits) {
// 			tmp := string(path)
// 			s = append(s, tmp)
// 			return
// 		}

// 		letters := letterMap[digits[index]-'0']
// 		for i := 0; i < len(letters); i++ {
// 			path = append(path, letters[i])
// 			backtracing(index + 1)
// 			path = path[:len(path)-1]
// 		}

// 	}
// 	backtracing(0)
// 	return s

// }

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ls := len(digits)
	lm := []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	var res []string
	var bt func(int, []byte)
	bt = func(start int, path []byte) {
		if ls == start {
			res = append(res, string(append([]byte{}, path...)))
			return
		}
		charsInMap := []byte(lm[digits[start]-'0'])
		for _, c := range charsInMap {
			path = append(path, c)
			bt(start+1, path)
			path = path[:len(path)-1]
		}
	}
	bt(0, []byte{})
	return res
}

// @lc code=end
