package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder3(root *TreeNode) [][]int {
	dqueueLevel := make([][]*TreeNode, 0)
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	currentNode := root
	dqueueLevel = append(dqueueLevel, []*TreeNode{currentNode})
	currentLevel := dqueueLevel[0]
	flag := false
	for len(dqueueLevel) > 0 {
		currentLevel = dqueueLevel[0]
		dqueueLevel = dqueueLevel[1:]
		var tempLevel []*TreeNode
		var tempLevelVal []int
		for _, n := range currentLevel {
			tempLevelVal = append(tempLevelVal, n.Val)
			if n.Left != nil {
				tempLevel = append(tempLevel, n.Left)
			}
			if n.Right != nil {
				tempLevel = append(tempLevel, n.Right)
			}
		}
		if len(tempLevel) > 0 {
			dqueueLevel = append(dqueueLevel, tempLevel)
		}
		if len(tempLevelVal) > 0 {
			if flag {
				reverse(tempLevelVal)
			}
			ret = append(ret, tempLevelVal)
			flag = !flag
		}
	}
	return ret
}

func reverse(qRes []int) {
	l, r := 0, len(qRes)-1
	for l <= r {
		qRes[l], qRes[r] = qRes[r], qRes[l]
		l++
		r--
	}
}
