package main

import (
	"fmt"
)

type TreeNodeWithParent struct {
	Data   int
	Left   *TreeNodeWithParent
	Right  *TreeNodeWithParent
	Parent *TreeNodeWithParent
}

type Tree struct {
	root *TreeNodeWithParent
}

func (t *Tree) Add(data int) *TreeNodeWithParent {
	var queue []*TreeNodeWithParent
	newNode := &TreeNodeWithParent{Data: data}
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
	return newNode
}

func succeedNode(n *TreeNodeWithParent) *TreeNodeWithParent {
	if n == nil {
		return nil
	}
	var cur *TreeNodeWithParent
	if n.Right != nil {
		cur = n.Right
		for cur.Left != nil {
			cur = cur.Left
		}
	} else if n.Parent != nil {
		cur = n
		parent := n.Parent
		for parent != nil && parent.Left != cur {
			cur = parent
			parent = parent.Parent
		}
		cur = parent
	}

	return cur
}

func (t *Tree) inorderTraverseRecursion(node *TreeNodeWithParent) {
	if node == nil {
		return
	} else {
		t.inorderTraverseRecursion(node.Left)
		fmt.Print(node.Data, " ")
		t.inorderTraverseRecursion(node.Right)

	}
}

func main() {

	// 			 3
	// 			/ \
	// 		   9  20
	// 			 /  \
	// 			15   7
	tree := &Tree{}
	tree.Add(3)
	tree.Add(9)
	tree.Add(20)
	n := tree.Add(15)
	tree.Add(7)

	tree.inorderTraverseRecursion(tree.root)
	fmt.Println()
	s := succeedNode(n)
	fmt.Println(s)
}
