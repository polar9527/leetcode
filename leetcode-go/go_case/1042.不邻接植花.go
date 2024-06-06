package go_case

/*
 * @lc app=leetcode.cn id=1042 lang=golang
 *
 * [1042] 不邻接植花
 *
 * https://leetcode.cn/problems/flower-planting-with-no-adjacent/description/
 *
 * algorithms
 * Medium (61.36%)
 * Likes:    238
 * Dislikes: 0
 * Total Accepted:    33.3K
 * Total Submissions: 54.3K
 * Testcase Example:  '3\n[[1,2],[2,3],[3,1]]'
 *
 * 有 n 个花园，按从 1 到 n 标记。另有数组 paths ，其中 paths[i] = [xi, yi] 描述了花园 xi 到花园 yi
 * 的双向路径。在每个花园中，你打算种下四种花之一。
 *
 * 另外，所有花园 最多 有 3 条路径可以进入或离开.
 *
 * 你需要为每个花园选择一种花，使得通过路径相连的任何两个花园中的花的种类互不相同。
 *
 * 以数组形式返回 任一 可行的方案作为答案 answer，其中 answer[i] 为在第 (i+1) 个花园中种植的花的种类。花的种类用
 * 1、2、3、4 表示。保证存在答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3, paths = [[1,2],[2,3],[3,1]]
 * 输出：[1,2,3]
 * 解释：
 * 花园 1 和 2 花的种类不同。
 * 花园 2 和 3 花的种类不同。
 * 花园 3 和 1 花的种类不同。
 * 因此，[1,2,3] 是一个满足题意的答案。其他满足题意的答案有 [1,2,4]、[1,4,2] 和 [3,2,1]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 4, paths = [[1,2],[3,4]]
 * 输出：[1,2,1,2]
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 4, paths = [[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]
 * 输出：[1,2,3,4]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^4
 * 0 <= paths.length <= 2 * 10^4
 * paths[i].length == 2
 * 1 <= xi, yi <= n
 * xi != yi
 * 每个花园 最多 有 3 条路径可以进入或离开
 *
 *
 */

// @lc code=start
func gardenNoAdj(n int, paths [][]int) []int {
	pathsMap := make([][]int, n)
	for i := 0; i < n; i++ {
		pathsMap[i] = []int{}
	}

	for _, path := range paths {
		pathsMap[path[0]-1] = append(pathsMap[path[0]-1], path[1]-1)
		pathsMap[path[1]-1] = append(pathsMap[path[1]-1], path[0]-1)
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		// 遍历 花园i 的 所有临接花园 pathsMap[i]
		used := make([]bool, 5)
		for _, vertex := range pathsMap[i] {
			// 该临接花园的花不能使用，登记到used字典中
			used[res[vertex]] = true
		}
		// 遍历 1-4 种花，找到第一个未被使用的颜色
		for j := 1; j <= 4; j++ {
			if used[j] == false {
				res[i] = j
				break
			}
		}
	}
	return res
}

// @lc code=end
