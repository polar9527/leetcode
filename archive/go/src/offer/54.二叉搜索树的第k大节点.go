package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var count int
	var result int
	kthLargestCore(root, &count, &result, k)
	return result
}

func kthLargestCore(node *TreeNode, count, result *int, k int) {
	if node == nil {
		return
	}

	kthLargestCore(node.Right, count, result, k)

	*count++
	if *count == k {
		*result = node.Val
	}
	kthLargestCore(node.Left, count, result, k)
}
