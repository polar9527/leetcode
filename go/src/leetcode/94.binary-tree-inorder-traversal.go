package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

	ret := []int{}

	s := []*TreeNode{}

	curNode := root
	for curNode != nil || len(s) != 0 {
		for curNode != nil {
			s = append(s, curNode)
			curNode = curNode.Left
		}

		curNode = s[len(s)-1]
		ret = append(ret, curNode.Val)
		s = s[:len(s)-1]
		curNode = curNode.Right

	}

	return ret
}
