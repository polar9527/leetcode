package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {

	ret := false
	if A != nil && B != nil {
		if A.Val == B.Val {
			ret = hasTree(A.Left, B.Left) && hasTree(A.Right, B.Right)
		}
		if !ret {
			ret = isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
		}
	}

	return ret
}

func hasTree(A *TreeNode, B *TreeNode) bool {

	if B == nil {
		return true
	}

	if A == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}

	return hasTree(A.Left, B.Left) && hasTree(A.Right, B.Right)
}
