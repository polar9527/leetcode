/*
 * @lc app=leetcode.cn id=732 lang=golang
 *
 * [732] 我的日程安排表 III
 *
 * https://leetcode.cn/problems/my-calendar-iii/description/
 *
 * algorithms
 * Hard (71.32%)
 * Likes:    205
 * Dislikes: 0
 * Total Accepted:    25.1K
 * Total Submissions: 35.2K
 * Testcase Example:  '["MyCalendarThree","book","book","book","book","book","book"]\n' +
  '[[],[10,20],[50,60],[10,40],[5,15],[5,10],[25,55]]'
 *
 * 当 k 个日程安排有一些时间上的交叉时（例如 k 个日程安排都在同一时间内），就会产生 k 次预订。
 *
 * 给你一些日程安排 [start, end) ，请你在每个日程安排添加后，返回一个整数 k ，表示所有先前日程安排会产生的最大 k 次预订。
 *
 * 实现一个 MyCalendarThree 类来存放你的日程安排，你可以一直添加新的日程安排。
 *
 *
 * MyCalendarThree() 初始化对象。
 * int book(int start, int end) 返回一个整数 k ，表示日历中存在的 k 次预订的最大值。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["MyCalendarThree", "book", "book", "book", "book", "book", "book"]
 * [[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
 * 输出：
 * [null, 1, 1, 2, 3, 3, 3]
 *
 * 解释：
 * MyCalendarThree myCalendarThree = new MyCalendarThree();
 * myCalendarThree.book(10, 20); // 返回 1 ，第一个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
 * myCalendarThree.book(50, 60); // 返回 1 ，第二个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
 * myCalendarThree.book(10, 40); // 返回 2 ，第三个日程安排 [10, 40) 与第一个日程安排相交，所以最大 k
 * 次预订是 2 次预订。
 * myCalendarThree.book(5, 15); // 返回 3 ，剩下的日程安排的最大 k 次预订是 3 次预订。
 * myCalendarThree.book(5, 10); // 返回 3
 * myCalendarThree.book(25, 55); // 返回 3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * 每个测试用例，调用 book 函数最多不超过 400次
 *
 *
*/

// @lc code=start
// 差分数组
// type MyCalendarThree struct {
//     *redblacktree.Tree
// }

// func Constructor() MyCalendarThree {
//     return MyCalendarThree{redblacktree.NewWithIntComparator()}
// }

// func (t MyCalendarThree) add(x, delta int) {
//     if val, ok := t.Get(x); ok {
//         delta += val.(int)
//     }
//     t.Put(x, delta)
// }

// func (t MyCalendarThree) Book(start, end int) (ans int) {
//     t.add(start, 1)
//     t.add(end, -1)

//     maxBook := 0
//     for it := t.Iterator(); it.Next(); {
//         maxBook += it.Value().(int)
//         if maxBook > ans {
//             ans = maxBook
//         }
//     }
//     return
// }

type pair struct{ num, lazy int }

type MyCalendarThree map[int]pair

func Constructor() MyCalendarThree {
	return MyCalendarThree{}
}

func (t MyCalendarThree) update(start, end, l, r, idx int) {
	if r < start || end < l {
		return
	}
	if start <= l && r <= end {
		p := t[idx]
		p.num++
		p.lazy++
		t[idx] = p
	} else {
		mid := (l + r) / 2
		t.update(start, end, l, mid, idx*2)
		t.update(start, end, mid+1, r, idx*2+1)
		p := t[idx]
		p.num = p.lazy + max(t[idx*2].num, t[idx*2+1].num)
		t[idx] = p
	}
}

func (t MyCalendarThree) Book(start, end int) int {
	t.update(start, end-1, 0, 1e9, 1)
	return t[1].num
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
// @lc code=end

