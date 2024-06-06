package go_case

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (71.64%)
 * Likes:    2308
 * Dislikes: 0
 * Total Accepted:    655.6K
 * Total Submissions: 914.9K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder
 * 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * 输出: [3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入: preorder = [-1], inorder = [-1]
 * 输出: [-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= preorder.length <= 3000
 * inorder.length == preorder.length
 * -3000 <= preorder[i], inorder[i] <= 3000
 * preorder 和 inorder 均 无重复 元素
 * inorder 均出现在 preorder
 * preorder 保证 为二叉树的前序遍历序列
 * inorder 保证 为二叉树的中序遍历序列
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
var (
	hash map[int]int
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	hash = make(map[int]int, len(inorder))
	for i, v := range inorder {
		hash[v] = i
	}
	root := build(preorder, inorder, 0, 0, len(inorder)-1) // l, r 表示构造的树在中序遍历数组中的范围
	return root
}
func build(pre []int, in []int, root int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	rootVal := pre[root]   // 找到本次构造的树的根节点
	index := hash[rootVal] // 根节点在中序数组中的位置
	node := &TreeNode{Val: rootVal}

	leftRoot := root + 1
	// leftInorder range [l, index)
	node.Left = build(pre, in, leftRoot, l, index-1)

	// root + len(leftInorder)
	// len(leftInorder) = index - l + 1
	rightRoot := root + (index - l) + 1
	// rightInorder range [l, index-1)
	node.Right = build(pre, in, rightRoot, index+1, r)
	return node
}

// @lc code=end
