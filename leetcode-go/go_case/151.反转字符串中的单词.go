/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 *
 * https://leetcode.cn/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (55.64%)
 * Likes:    1274
 * Dislikes: 0
 * Total Accepted:    706.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '"the sky is blue"'
 *
 * 给你一个字符串 s ，请你反转字符串中 单词 的顺序。
 *
 * 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
 *
 * 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
 *
 * 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "the sky is blue"
 * 输出："blue is sky the"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "  hello world  "
 * 输出："world hello"
 * 解释：反转后的字符串中不能存在前导空格和尾随空格。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "a good   example"
 * 输出："example good a"
 * 解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 包含英文大小写字母、数字和空格 ' '
 * s 中 至少存在一个 单词
 *
 *
 *
 *
 *
 *
 *
 * 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。
 *
 */

// @lc code=start
// func reverseWords(s string) string {
// 	sb := []byte(s)

// 	slow := 0
// 	for fast := 0; fast < len(sb); fast++ {
// 		// sb[fast] 为空格， fast 继续往后走
// 		if sb[fast] != byte(' ') {
// 			// 补上单词之间的一个空格
// 			if slow != 0 {
// 				sb[slow] = byte(' ')
// 				slow++
// 			}
// 			// 把单词整体往前挪动
// 			for fast < len(sb) && sb[fast] != byte(' ') {
// 				sb[slow] = sb[fast]
// 				fast++
// 				slow++
// 			}
// 		}
// 	}
// 	sb = sb[:slow]

// 	reverse := func(sb []byte) []byte {
// 		l, r := 0, len(sb)-1
// 		for l < r {
// 			sb[l] ^= sb[r]
// 			sb[r] ^= sb[l]
// 			sb[l] ^= sb[r]
// 			l++
// 			r--
// 		}
// 		return sb
// 	}

// 	reverse(sb)

// 	head := 0
// 	for i, c := range sb {
// 		if c == byte(' ') {
// 			reverse(sb[head:i])
// 			head = i + 1
// 		}
// 		// 处理最后一个单词
// 		if i == len(sb)-1 {
// 			reverse(sb[head : i+1])
// 		}

// 	}
// 	return string(sb)
// }

func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//2.反转整个字符串
	reverse(b)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(b[i:j])
		i = j
		i++
	}
	return string(b)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// @lc code=end

