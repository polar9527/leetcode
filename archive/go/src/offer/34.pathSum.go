package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	// result := make([][]int, 0)
	path := []int{}
	if root == nil {
		return result
	}

	printPath(&result, &path, sum, 0, root)
	return result
}

func printPath(ret *[][]int, path *[]int, sum, currentSum int, currentNode *TreeNode) {

	currentSum += currentNode.Val
	*path = append(*path, currentNode.Val)
	if currentSum == sum && currentNode.Left == nil && currentNode.Right == nil {
		var temp []int
		for _, v := range *path {
			temp = append(temp, v)
		}
		*ret = append(*ret, temp)
	}

	if currentNode.Left != nil {
		printPath(ret, path, sum, currentSum, currentNode.Left)
	}
	if currentNode.Right != nil {
		printPath(ret, path, sum, currentSum, currentNode.Right)
	}

	*path = (*path)[:len(*path)-1]
}

// func main() {
// 	root := TreeNode{5, nil, nil}
// 	root.Left = &TreeNode{4, nil, nil}
// 	root.Left.Left = &TreeNode{11, nil, nil}
// 	root.Left.Left.Left = &TreeNode{7, nil, nil}
// 	root.Left.Left.Right = &TreeNode{2, nil, nil}
// 	root.Right = &TreeNode{8, nil, nil}
// 	root.Right.Left = &TreeNode{13, nil, nil}
// 	root.Right.Right = &TreeNode{4, nil, nil}
// 	root.Right.Right.Left = &TreeNode{5, nil, nil}
// 	root.Right.Right.Right = &TreeNode{1, nil, nil}
//
// 	ret := pathSum(&root, 22)
// 	fmt.Println(ret)
//
// }
