package go_case

/*
 * @lc app=leetcode.cn id=1462 lang=golang
 *
 * [1462] 课程表 IV
 *
 * https://leetcode.cn/problems/course-schedule-iv/description/
 *
 * algorithms
 * Medium (51.70%)
 * Likes:    216
 * Dislikes: 0
 * Total Accepted:    35.9K
 * Total Submissions: 69.4K
 * Testcase Example:  '2\n[[1,0]]\n[[0,1],[1,0]]'
 *
 * 你总共需要上 numCourses 门课，课程编号依次为 0 到 numCourses-1 。你会得到一个数组 prerequisite ，其中
 * prerequisites[i] = [ai, bi] 表示如果你想选 bi 课程，你 必须 先选 ai 课程。
 *
 *
 * 有的课会有直接的先修课程，比如如果想上课程 1 ，你必须先上课程 0 ，那么会以 [0,1] 数对的形式给出先修课程数对。
 *
 *
 * 先决条件也可以是 间接 的。如果课程 a 是课程 b 的先决条件，课程 b 是课程 c 的先决条件，那么课程 a 就是课程 c 的先决条件。
 *
 * 你也得到一个数组 queries ，其中 queries[j] = [uj, vj]。对于第 j 个查询，您应该回答课程 uj 是否是课程 vj
 * 的先决条件。
 *
 * 返回一个布尔数组 answer ，其中 answer[j] 是第 j 个查询的答案。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]], queries = [[0,1],[1,0]]
 * 输出：[false,true]
 * 解释：[1, 0] 数对表示在你上课程 0 之前必须先上课程 1。
 * 课程 0 不是课程 1 的先修课程，但课程 1 是课程 0 的先修课程。
 *
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 2, prerequisites = [], queries = [[1,0],[0,1]]
 * 输出：[false,false]
 * 解释：没有先修课程对，所以每门课程之间是独立的。
 *
 *
 * 示例 3：
 *
 *
 *
 *
 * 输入：numCourses = 3, prerequisites = [[1,2],[1,0],[2,0]], queries =
 * [[1,0],[1,2]]
 * 输出：[true,true]
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 2 <= numCourses <= 100
 * 0 <= prerequisites.length <= (numCourses * (numCourses - 1) / 2)
 * prerequisites[i].length == 2
 * 0 <= ai, bi <= numCourses - 1
 * ai != bi
 * 每一对 [ai, bi] 都 不同
 * 先修课程图中没有环。
 * 1 <= queries.length <= 10^4
 * 0 <= ui, vi <= numCourses - 1
 * ui != vi
 *
 *
 */

// @lc code=start
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {

	adjm := make([][]bool, numCourses)
	for i := range adjm {
		adjm[i] = make([]bool, numCourses)
	}

	n := len(prerequisites)
	for i := 0; i < n; i++ {
		from, to := prerequisites[i][0], prerequisites[i][1]
		adjm[from][to] = true
	}

	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				if adjm[i][k] && adjm[k][j] {
					adjm[i][j] = true
				}
			}
		}
	}

	res := make([]bool, len(queries))
	for i := range queries {
		from, to := queries[i][0], queries[i][1]
		if adjm[from][to] {
			res[i] = true
		}
	}
	return res
}

// @lc code=end
