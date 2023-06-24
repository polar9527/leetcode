package main

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {

	var depth int
	return isBalancedTree(root, &depth)
}

func isBalancedTree(node *TreeNode, depth *int) bool {
	if node == nil {
		*depth = 0
		return true
	}
	var leftDepth int
	var rightDepth int
	l := isBalancedTree(node.Left, &leftDepth)
	r := isBalancedTree(node.Right, &rightDepth)

	if l && r {
		diff := leftDepth - rightDepth
		if diff <= 1 && diff >= -1 {
			*depth = int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
			return true
		}
	}
	return false
}
