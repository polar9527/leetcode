/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode.cn/problems/course-schedule/description/
 *
 * algorithms
 * Medium (55.05%)
 * Likes:    2130
 * Dislikes: 0
 * Total Accepted:    539.9K
 * Total Submissions: 977.7K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 *
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi]
 * ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
 *
 *
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 *
 *
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= 5000
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * prerequisites[i] 中的所有课程对 互不相同
 *
 *
 */

// @lc code=start
// func canFinish(numCourses int, prerequisites [][]int) bool {

// 	inDegree := make(map[int]int)
// 	dependencies := make(map[int][]int)
// 	dn := len(prerequisites)
// 	for i := 0; i < dn; i++ {
// 		to, from := prerequisites[i][0], prerequisites[i][1]
// 		inDegree[to]++
// 		dependencies[from] = append(dependencies[from], to)
// 	}

// 	que := make([]int, 0)
// 	for i := 0; i < numCourses; i++ {
// 		if inDegree[i] == 0 {
// 			que = append(que, i)
// 		}
// 	}

// 	count := 0
// 	for len(que) > 0 {
// 		cur := que[0]
// 		que = que[1:]
// 		count++

// 		files := dependencies[cur]
// 		if len(files) != 0 {
// 			for _, f := range files {
// 				inDegree[f]--
// 				if inDegree[f] == 0 {
// 					que = append(que, f)
// 				}
// 			}
// 		}
// 	}

// 	return count == numCourses
// }

// 拓扑排序 优化版本，不用字典
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegrees := make([]int, numCourses)
	dependences := make([][]int, numCourses)
	for _, prerequisit := range prerequisites {
		// from -> to
		to, from := prerequisit[0], prerequisit[1]
		inDegrees[to]++
		// from -> to1
		// 	 	-> to2
		// 		-> ...
		dependences[from] = append(dependences[from], to)
	}

	queue := []int{}
	for to, degree := range inDegrees {
		if degree == 0 {
			queue = append(queue, to)
		}
	}
	count := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		for _, next := range dependences[cur] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return count == numCourses
}

// 环检测
// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	g := make([][]int, numCourses)
// 	for _, p := range prerequisites {
// 		g[p[1]] = append(g[p[1]], p[0])
// 	}

// 	colors := make([]int, numCourses)
// 	var dfs func(int) bool
// 	dfs = func(x int) bool {
// 		colors[x] = 1 // x 正在访问中
// 		for _, y := range g[x] {
// 			if colors[y] == 1 || colors[y] == 0 && dfs(y) {
// 				return true // 找到了环
// 			}
// 		}
// 		colors[x] = 2 // x 完全访问完毕
// 		return false  // 没有找到环
// 	}

// 	for i, c := range colors {
// 		if c == 0 && dfs(i) {
// 			return false // 有环
// 		}
// 	}
// 	return true // 没有环
// }

// @lc code=end

