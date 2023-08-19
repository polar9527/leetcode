package go_case

/*
 * @lc app=leetcode.cn id=731 lang=golang
 *
 * [731] 我的日程安排表 II
 *
 * https://leetcode.cn/problems/my-calendar-ii/description/
 *
 * algorithms
 * Medium (62.68%)
 * Likes:    223
 * Dislikes: 0
 * Total Accepted:    28.6K
 * Total Submissions: 45.6K
 * Testcase Example:  '["MyCalendarTwo","book","book","book","book","book","book"]\n' +
  '[[],[10,20],[50,60],[10,40],[5,15],[5,10],[25,55]]'
 *
 * 实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内不会导致三重预订时，则可以存储这个新的日程安排。
 *
 * MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end
 * 时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。
 *
 * 当三个日程安排有一些时间上的交叉时（例如三个日程安排都在同一时间内），就会产生三重预订。
 *
 * 每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致三重预订，返回 true。否则，返回 false
 * 并且不要将该日程安排添加到日历中。
 *
 * 请按照以下步骤调用MyCalendar 类: MyCalendar cal = new MyCalendar();
 * MyCalendar.book(start, end)
 *
 *
 *
 * 示例：
 *
 * MyCalendar();
 * MyCalendar.book(10, 20); // returns true
 * MyCalendar.book(50, 60); // returns true
 * MyCalendar.book(10, 40); // returns true
 * MyCalendar.book(5, 15); // returns false
 * MyCalendar.book(5, 10); // returns true
 * MyCalendar.book(25, 55); // returns true
 * 解释：
 * 前两个日程安排可以添加至日历中。 第三个日程安排会导致双重预订，但可以添加至日历中。
 * 第四个日程安排活动（5,15）不能添加至日历中，因为它会导致三重预订。
 * 第五个日程安排（5,10）可以添加至日历中，因为它未使用已经双重预订的时间10。
 * 第六个日程安排（25,55）可以添加至日历中，因为时间 [25,40] 将和第三个日程安排双重预订；
 * 时间 [40,50] 将单独预订，时间 [50,55）将和第二个日程安排双重预订。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每个测试用例，调用 MyCalendar.book 函数最多不超过 1000次。
 * 调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。
 *
 *
*/

// @lc code=start
// type pair struct{ start, end int }
// type MyCalendarTwo struct{ booked, overlaps []pair }

// func Constructor() MyCalendarTwo {
//     return MyCalendarTwo{}
// }

// func (c *MyCalendarTwo) Book(start, end int) bool {
//     for _, p := range c.overlaps {
//         if p.start < end && start < p.end {
//             return false
//         }
//     }
//     for _, p := range c.booked {
//         if p.start < end && start < p.end {
//             c.overlaps = append(c.overlaps, pair{max(p.start, start), min(p.end, end)})
//         }
//     }
//     c.booked = append(c.booked, pair{start, end})
//     return true
// }

// func min(a, b int) int {
//     if a > b {
//         return b
//     }
//     return a
// }

// func max(a, b int) int {
//     if b > a {
//         return b
//     }
//     return a
// }

// 差分数组
// type MyCalendarTwo struct {
//     *redblacktree.Tree
// }

// func Constructor() MyCalendarTwo {
//     return MyCalendarTwo{redblacktree.NewWithIntComparator()}
// }

// func (c MyCalendarTwo) add(key, value int) {
//     if v, ok := c.Get(key); ok {
//         c.Put(key, v.(int)+value)
//     } else {
//         c.Put(key, value)
//     }
// }

// func (c MyCalendarTwo) Book(start, end int) bool {
//     c.add(start, 1)
//     c.add(end, -1)
//     maxBook := 0
//     it := c.Iterator()
//     for it.Next() {
//         maxBook += it.Value().(int)
//         if maxBook > 2 {
//             c.add(start, -1)
//             c.add(end, 1)
//             return false
//         }
//     }
//     return true
// }

type pair struct{ first, second int }
type MyCalendarTwo map[int]pair

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (tree MyCalendarTwo) update(start, end, val, l, r, idx int) {
	if r < start || end < l {
		return
	}
	if start <= l && r <= end {
		p := tree[idx]
		p.first += val
		p.second += val
		tree[idx] = p
		return
	}
	mid := (l + r) >> 1
	tree.update(start, end, val, l, mid, 2*idx)
	tree.update(start, end, val, mid+1, r, 2*idx+1)
	p := tree[idx]
	p.first = p.second + max(tree[2*idx].first, tree[2*idx+1].first)
	tree[idx] = p
}

func (tree MyCalendarTwo) Book(start, end int) bool {
	tree.update(start, end-1, 1, 0, 1e9, 1)
	if tree[1].first > 2 {
		tree.update(start, end-1, -1, 0, 1e9, 1)
		return false
	}
	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
