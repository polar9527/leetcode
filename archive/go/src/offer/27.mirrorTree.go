package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	tempNode := root.Left
	root.Left = root.Right
	root.Right = tempNode
	if root.Left != nil {
		mirrorTree(root.Left)
	}

	if root.Right != nil {
		mirrorTree(root.Right)
	}

	return root
}
