package go_case

import (
	"container/list"
	"math"
)

/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 *
 * https://leetcode.cn/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (66.42%)
 * Likes:    326
 * Dislikes: 0
 * Total Accepted:    126.3K
 * Total Submissions: 190.1K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
 *
 *
 *
 * 示例1：
 *
 *
 *
 *
 * 输入: root = [1,3,2,5,3,null,9]
 * 输出: [1,3,9]
 *
 *
 * 示例2：
 *
 *
 * 输入: root = [1,2,3]
 * 输出: [1,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 二叉树的节点个数的范围是 [0,10^4]
 * -2^31 <= Node.val <= 2^31 - 1
 *
 *
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	if root == nil {
		//防止为空
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	ans := make([]int, 0)

	// 层序遍历
	for queue.Len() > 0 {
		//保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		length := queue.Len()
		temp := math.MinInt64
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			// 比较当前层中的最大值和新遍历的元素大小，取两者中大值
			temp = max(temp, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans = append(ans, temp)
	}
	return ans
}

// @lc code=end
