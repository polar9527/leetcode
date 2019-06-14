/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    
    lDepth := maxDepth(root.Left)
    rDepth := maxDepth(root.Right)
    return max(lDepth, rDepth) + 1
}

func max(a int, b int) int {
    if a > b{
        return a
    }
    return b
}

