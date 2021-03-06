package main

import (
	"fmt"
)

type TreeNode struct {
	Data   int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) Add(data int) {
	var queue []*TreeNode
	newNode := &TreeNode{Data: data}
	if t.root == nil {
		t.root = newNode
	} else {
		queue = append(queue, t.root)
		for len(queue) != 0 {
			cur := queue[0]
			queue = append(queue[:0], queue[0+1:]...)
			// 往右树添加
			if data > cur.Data {
				if cur.Right == nil {
					cur.Right = newNode
					newNode.Parent = cur
				} else {
					queue = append(queue, cur.Right)
				}
				// 往左数添加
			} else {
				if cur.Left == nil {
					cur.Left = newNode
					newNode.Parent = cur
				} else {
					queue = append(queue, cur.Left)
				}
			}
		}
	}
}

func visitAlongLeftBranch(n *TreeNode, s *[]*TreeNode) {
	cur := n
	for cur != nil {
		fmt.Print(cur.Data, " ")
		*s = append(*s, cur.Right)
		cur = cur.Left
	}
}

func (t *Tree) preorderTraverseIteration(node *TreeNode) {
	var stack []*TreeNode
	cur := node
	for {
		visitAlongLeftBranch(cur, &stack)
		if len(stack) == 0 {
			break
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
}

func goAlongLeftBranch(node *TreeNode, s *[]*TreeNode) {
	cur := node
	for cur != nil {
		*s = append(*s, cur)
		cur = cur.Left
	}
}

func (t *Tree) inorderTraverseIteration(node *TreeNode) {
	var stack []*TreeNode
	cur := node

	for {
		goAlongLeftBranch(cur, &stack)
		if len(stack) == 0 {
			break
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(cur.Data, " ")
		cur = cur.Right
	}
}

// highest left visable leaf
func gotoHLVL(s *[]*TreeNode) {
	cur := (*s)[len(*s)-1]
	for ; cur != nil; cur = (*s)[len(*s)-1] {
		if cur.Left != nil {
			if cur.Right != nil {
				*s = append(*s, cur.Right)
			}
			*s = append(*s, cur.Left)
		} else {
			*s = append(*s, cur.Right)
		}
	}
	*s = (*s)[:len(*s)-1]
}

func (t *Tree) postorderTraverseIteration(node *TreeNode) {
	var stack []*TreeNode
	cur := node
	if node != nil {
		stack = append(stack, cur)
	}
	for len(stack) != 0 {
		if stack[len(stack)-1] != cur.Parent {
			gotoHLVL(&stack)
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(cur.Data, " ")
	}

}

func (t *Tree) preorderTraverseRecursion(node *TreeNode) {
	if node == nil {
		return
	} else {
		fmt.Print(node.Data, " ")
		t.preorderTraverseRecursion(node.Left)
		t.preorderTraverseRecursion(node.Right)
	}

}

func (t *Tree) inorderTraverseRecursion(node *TreeNode) {
	if node == nil {
		return
	} else {
		t.inorderTraverseRecursion(node.Left)
		fmt.Print(node.Data, " ")
		t.inorderTraverseRecursion(node.Right)

	}
}

func (t *Tree) postorderTraverseRecursion(node *TreeNode) {
	if node == nil {
		return
	} else {
		t.postorderTraverseRecursion(node.Left)
		t.postorderTraverseRecursion(node.Right)
		fmt.Print(node.Data, " ")
	}
}

func main() {
	tree := &Tree{}
	tree.Add(50)
	tree.Add(45)
	tree.Add(40)
	tree.Add(48)
	tree.Add(51)
	tree.Add(61)
	tree.Add(71)

	fmt.Println("preorderTraverse")
	fmt.Println("Recursion")
	tree.preorderTraverseRecursion(tree.root)
	fmt.Println()
	fmt.Println("Iteration")
	tree.preorderTraverseIteration(tree.root)

	fmt.Println("")

	fmt.Println("inorderTraverse")
	fmt.Println("Recursion")
	tree.inorderTraverseRecursion(tree.root)
	fmt.Println()
	fmt.Println("Iteration")
	tree.inorderTraverseIteration(tree.root)

	fmt.Println("")

	fmt.Println("postorderTraverse")
	fmt.Println("Recursion")
	tree.postorderTraverseRecursion(tree.root)
	fmt.Println()
	fmt.Println("Iteration")
	tree.postorderTraverseIteration(tree.root)

}
