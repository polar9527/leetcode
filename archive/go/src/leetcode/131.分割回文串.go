/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */
package leetcode

// @lc code=start
func partition(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	// 动态规划求出所有s[i:j]组合是否为回文串
	for i := n - 1; i >= 0; i-- {
		// if i >= j, f[i][j] == true
		// i < j, then
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	var ret [][]string
	var result []string
	// ret = [][]string{result, result}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			// 因为 result 是整个回溯过程（深度遍历）中的共享容器，所以要在此处将值拷贝一遍
			ret = append(ret, append([]string(nil), result...))
		}

		for j := i; j < n; j++ {
			if f[i][j] {
				result = append(result, s[i:j+1])
				dfs(j + 1)
				result = result[:len(result)-1]
			}
		}
	}
	dfs(0)
	return ret
}

func partitionMemory(s string) [][]string {
	n := len(s)
	f := make([][]int8, n)
	for i := range f {
		f[i] = make([]int8, n)
	}

	// 动态规划求出所有s[i:j]组合是否为回文串
	//  0 表示未被搜索过
	//  1 表示是Palindrome
	// -1 表示不是Palindrome
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if f[i][j] != 0 {
			return
		}
		f[i][j] = -1
		if s[i] == s[j] {
			f[i][j] = isPalindrome(i+1, j-1)
		}
	}

	var ret [][]string
	var result []string
	// ret = [][]string{result, result}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			// 因为 result 是整个回溯过程（深度遍历）中的共享容器，所以要在此处将值拷贝一遍
			ret = append(ret, append([]string(nil), result...))
		}

		for j := i; j < n; j++ {
			if isPalindrome(i, j) {
				result = append(result, s[i:j+1])
				dfs(j + 1)
				result = result[:len(result)-1]
			}
		}
	}
	dfs(0)
	return ret
}

// @lc code=end
