/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (75.18%)
 * Likes:    1833
 * Dislikes: 0
 * Total Accepted:    637.7K
 * Total Submissions: 843.8K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给你二叉树的根结点 root ，请你将它展开为一个单链表：
 *
 *
 * 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
 * 展开后的单链表应该与二叉树 先序遍历 顺序相同。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,5,3,4,null,6]
 * 输出：[1,null,2,null,3,null,4,null,5,null,6]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中结点数在范围 [0, 2000] 内
 * -100
 *
 *
 *
 *
 * 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
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
// func flatten(root *TreeNode) {

// 	var head *TreeNode
// 	var dfs func(*TreeNode)
// 	dfs = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		dfs(node.Right)
// 		dfs(node.Left)

// 		node.Left = nil
// 		node.Right = head
// 		head = node
// 	}
// 	dfs(root)
// }

func flatten(root *TreeNode) {

	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		leftTail := dfs(node.Left)
		rightTail := dfs(node.Right)

		// 首先 将 左子树的链表 插入到 父节点和右子树链表之间
		if leftTail != nil {
			leftTail.Right = node.Right
			node.Right = node.Left
			node.Left = nil
		}
		if rightTail != nil {
			return rightTail
		}
		// 此时 rightTail == nil
		// 所以 链表尾 是 leftTail
		if leftTail != nil {
			return leftTail
		}
		// 此时 rightTail 和 leftTail 都是 nil
		return node
	}
	dfs(root)
}

// @lc code=end

