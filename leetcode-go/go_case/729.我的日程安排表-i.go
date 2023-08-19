package go_case

import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

/*
 * @lc app=leetcode.cn id=729 lang=golang
 *
 * [729] 我的日程安排表 I
 *
 * https://leetcode.cn/problems/my-calendar-i/description/
 *
 * algorithms
 * Medium (58.31%)
 * Likes:    265
 * Dislikes: 0
 * Total Accepted:    52.6K
 * Total Submissions: 90.3K
 * Testcase Example:  '["MyCalendar","book","book","book"]\n[[],[10,20],[15,25],[20,30]]'
 *
 * 实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。
 *
 * 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。
 *
 * 日程可以用一对整数 start 和 end 表示，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x <
 * end 。
 *
 * 实现 MyCalendar 类：
 *
 *
 * MyCalendar() 初始化日历对象。
 * boolean book(int start, int end) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true 。否则，返回
 * false 并且不要将该日程安排添加到日历中。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["MyCalendar", "book", "book", "book"]
 * [[], [10, 20], [15, 25], [20, 30]]
 * 输出：
 * [null, true, false, true]
 *
 * 解释：
 * MyCalendar myCalendar = new MyCalendar();
 * myCalendar.book(10, 20); // return True
 * myCalendar.book(15, 25); // return False ，这个日程安排不能添加到日历中，因为时间 15
 * 已经被另一个日程安排预订了。
 * myCalendar.book(20, 30); // return True ，这个日程安排可以添加到日历中，因为第一个日程安排预订的每个时间都小于
 * 20 ，且不包含时间 20 。
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= start < end <= 10^9
 * 每个测试用例，调用 book 方法的次数最多不超过 1000 次。
 *
 *
 */

// @lc code=start
type MyCalendar struct {
	*redblacktree.Tree
}

func Constructor() MyCalendar {
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt32, nil) // 哨兵，简化代码
	return MyCalendar{t}
}

func (c MyCalendar) Book(start, end int) bool {
	node, _ := c.Ceiling(end)
	it := c.IteratorAt(node)
	if !it.Prev() || it.Value().(int) <= start {
		c.Put(start, end)
		return true
	}
	return false
}

// brute force
// type pair struct{ start, end int }
// type MyCalendar []pair

// func Constructor() MyCalendar {
//     return MyCalendar{}
// }

// func (c *MyCalendar) Book(start, end int) bool {
//     for _, p := range *c {
//         if p.start < end && start < p.end {
//             return false
//         }
//     }
//     *c = append(*c, pair{start, end})
//     return true
// }

// segment tree
// type MyCalendar struct {
//     tree, lazy map[int]bool
// }

// func Constructor() MyCalendar {
//     return MyCalendar{map[int]bool{}, map[int]bool{}}
// }

// func (c MyCalendar) query(start, end, l, r, idx int) bool {
//     if r < start || end < l {
//         return false
//     }
//     if c.lazy[idx] { // 如果该区间已被预订，则直接返回
//         return true
//     }
//     if start <= l && r <= end {
//         return c.tree[idx]
//     }
//     mid := (l + r) >> 1
//     return c.query(start, end, l, mid, 2*idx) ||
//         c.query(start, end, mid+1, r, 2*idx+1)
// }

// func (c MyCalendar) update(start, end, l, r, idx int) {
//     if r < start || end < l {
//         return
//     }
//     if start <= l && r <= end {
//         c.tree[idx] = true
//         c.lazy[idx] = true
//     } else {
//         mid := (l + r) >> 1
//         c.update(start, end, l, mid, 2*idx)
//         c.update(start, end, mid+1, r, 2*idx+1)
//         c.tree[idx] = true
//         if c.lazy[2*idx] && c.lazy[2*idx+1] {
//             c.lazy[idx] = true
//         }
//     }
// }

// func (c MyCalendar) Book(start, end int) bool {
//     if c.query(start, end-1, 0, 1e9, 1) {
//         return false
//     }
//     c.update(start, end-1, 0, 1e9, 1)
//     return true
// }

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
