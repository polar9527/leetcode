package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder1(root *TreeNode) []int {

	dqueueNode := make([]*TreeNode, 0)
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	currentNode := root
	dqueueNode = append(dqueueNode, currentNode)
	for len(dqueueNode) > 0 {
		ret = append(ret, dqueueNode[0].Val)
		currentNode = dqueueNode[0]
		dqueueNode = dqueueNode[1:]
		if currentNode.Left != nil {
			dqueueNode = append(dqueueNode, currentNode.Left)
		}
		if currentNode.Right != nil {
			dqueueNode = append(dqueueNode, currentNode.Right)
		}
	}
	return ret
}
