/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (45.74%)
 * Likes:    2696
 * Dislikes: 78
 * Total Accepted:    310.6K
 * Total Submissions: 674.4K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given preorder and inorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * preorder = [3,9,20,15,7]
 * inorder = [9,3,15,20,7]
 *
 * Return the following binary tree:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 */
package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	var rootIndex int
	for i, v := range inorder {
		if v == preorder[0] {
			rootIndex = i
			break
		}
	}
	lInorder := inorder[:rootIndex]
	if len(lInorder) > 0 {
		lPreorder := preorder[1 : len(lInorder)+1]
		root.Left = buildTree(lPreorder, lInorder)
	}

	if rootIndex < len(inorder)-1 {
		rInorder := inorder[rootIndex+1:]
		rPreorder := preorder[len(lInorder)+1:]
		root.Right = buildTree(rPreorder, rInorder)
	}

	return root
}

// @lc code=end
