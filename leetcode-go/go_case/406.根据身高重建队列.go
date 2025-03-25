/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 *
 * https://leetcode.cn/problems/queue-reconstruction-by-height/description/
 *
 * algorithms
 * Medium (76.91%)
 * Likes:    1903
 * Dislikes: 0
 * Total Accepted:    333.3K
 * Total Submissions: 433.3K
 * Testcase Example:  '[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]'
 *
 * 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第
 * i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
 *
 * 请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj]
 * 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
 * 输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
 * 解释：
 * 编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
 * 编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
 * 编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
 * 编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
 * 编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
 * 编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
 * 因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
 * 输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0 i
 * 0 i < people.length
 * 题目数据确保队列可以被重建
 *
 *
 */

// @lc code=start
// func reconstructQueue(people [][]int) [][]int {
// 	sort.Slice(people, func(i, j int) bool {
// 		// 当身高hi相同，按照ki 从小到大排序
// 		if people[i][0] == people[j][0] {
// 			return people[i][1] < people[j][1]
// 		}
// 		// 按照身高从大到小排序
// 		return people[i][0] > people[j][0]
// 	})

//		l := list.New()
//		for i := 0; i < len(people); i++ {
//			position := people[i][1]
//			e := l.PushBack(people[i])
//			mark := l.Front()
//			for position > 0 {
//				position--
//				mark = mark.Next()
//			}
//			// move e before mark
//			l.MoveBefore(e, mark)
//		}
//		var res [][]int
//		for i := l.Front(); i != nil; i = i.Next() {
//			res = append(res, i.Value.([]int))
//		}
//		return res
//	}
func reconstructQueue(people [][]int) [][]int {
	// 先将身高从大到小排序，确定最大个子的相对位置
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] // 当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0] // 身高按照由大到小的顺序来排
	})

	// 再按照K进行插入排序，优先插入K小的
	for i, p := range people {
		// 空出一个位置
		copy(people[p[1]+1:i+1], people[p[1]:i])
		people[p[1]] = p
	}
	return people
}

// @lc code=end

