package go_case

import (
	"container/list"
	"sort"
)

/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 *
 * https://leetcode.com/problems/queue-reconstruction-by-height/description/
 *
 * algorithms
 * Medium (73.39%)
 * Likes:    7018
 * Dislikes: 712
 * Total Accepted:    298.9K
 * Total Submissions: 407.3K
 * Testcase Example:  '[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]'
 *
 * You are given an array of people, people, which are the attributes of some
 * people in a queue (not necessarily in order). Each people[i] = [hi, ki]
 * represents the i^th person of height hi with exactly ki other people in
 * front who have a height greater than or equal to hi.
 *
 * Reconstruct and return the queue that is represented by the input array
 * people. The returned queue should be formatted as an array queue, where
 * queue[j] = [hj, kj] is the attributes of the j^th person in the queue
 * (queue[0] is the person at the front of the queue).
 *
 *
 * Example 1:
 *
 *
 * Input: people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
 * Output: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
 * Explanation:
 * Person 0 has height 5 with no other people taller or the same height in
 * front.
 * Person 1 has height 7 with no other people taller or the same height in
 * front.
 * Person 2 has height 5 with two persons taller or the same height in front,
 * which is person 0 and 1.
 * Person 3 has height 6 with one person taller or the same height in front,
 * which is person 1.
 * Person 4 has height 4 with four people taller or the same height in front,
 * which are people 0, 1, 2, and 3.
 * Person 5 has height 7 with one person taller or the same height in front,
 * which is person 1.
 * Hence [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] is the reconstructed queue.
 *
 *
 * Example 2:
 *
 *
 * Input: people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
 * Output: [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= people.length <= 2000
 * 0 <= hi <= 10^6
 * 0 <= ki < people.length
 * It is guaranteed that the queue can be reconstructed.
 *
 *
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] //当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0]
	})
	l := list.New() //创建链表
	for i := 0; i < len(people); i++ {
		position := people[i][1]
		mark := l.PushBack(people[i]) //插入元素
		e := l.Front()
		for position != 0 { //获取相对位置
			position--
			e = e.Next()
		}
		l.MoveBefore(mark, e) //移动位置

	}
	res := [][]int{}
	for e := l.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.([]int))
	}
	return res
}

// func reconstructQueue(people [][]int) [][]int {
//     // 先将身高从大到小排序，确定最大个子的相对位置
//     sort.Slice(people, func(i, j int) bool {
//         if people[i][0] == people[j][0] {
//             return people[i][1] < people[j][1]   // 当身高相同时，将K按照从小到大排序
//         }
//         return people[i][0] > people[j][0]     // 身高按照由大到小的顺序来排
//     })

//     // 再按照K进行插入排序，优先插入K小的
// 	for i, p := range people {
// 		copy(people[p[1]+1 : i+1], people[p[1] : i+1])  // 空出一个位置
// 		people[p[1]] = p
// 	}
// 	return people
// }

// @lc code=end
