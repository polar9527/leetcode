package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	buildTree(pre, in)
}

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
