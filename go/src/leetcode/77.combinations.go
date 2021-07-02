/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (52.98%)
 * Likes:    1285
 * Dislikes: 63
 * Total Accepted:    271.6K
 * Total Submissions: 512K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 *
 * Example:
 *
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 *
 */
package leetcode

// @lc code=start
func combine(n int, k int) [][]int {
	if n == 0 {
		return nil
	}
	if k > n {
		return nil
	}
	res := make([][]int, 0)

	path := make([]int, 0)
	combineCore(1, 1, n, k, path, &res)
	return res
}

func combineCore(start, depth, n, k int, path []int, res *[][]int) {
	if depth > k {
		var resultForAdd = make([]int, k)
		copy(resultForAdd, path)
		*res = append(*res, resultForAdd)
		return
	}

	for i := start; i <= n; i++ {
		path = append(path, i)
		combineCore(i+1, depth+1, n, k, path, res)
		path = path[:len(path)-1]

	}

}

// @lc code=end

// func main(){
// 	res := combine(4,2)
// 	fmt.Println(res)
// }
