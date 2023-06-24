package bintree

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
			//  取出queue中的首元素
			cur := queue[0]
			// 删除queue中的首元素，而且是就地操作，仔细看下面的写法
			queue = append(queue[:0], queue[1:]...)
			// 往右树添加
			if newNode.Data > cur.Data {
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

func visit(data int) {
	fmt.Print(data, " ")
}

func visitAlongLeftBranch(n *TreeNode, s *[]*TreeNode) {
	cur := n
	for cur != nil {
		visit(cur.Data)
		// 右子树为空则不加入栈
		if cur.Right != nil {
			*s = append(*s, cur.Right)
		}
		// step forward
		cur = cur.Left
	}
}

func (t *Tree) preorderTraverseIteration(node *TreeNode) {
	var stack []*TreeNode
	cur := node
	for {
		// 自顶向下，依次访问当前节点的最左侧通路的各节点
		visitAlongLeftBranch(cur, &stack)
		if len(stack) == 0 {
			break
		}
		// 出栈
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
}

func goAlongLeftBranch(node *TreeNode, s *[]*TreeNode) {
	cur := node
	for cur != nil {
		// 将沿途的节点存入栈中，以便后续出栈后访问其右子树
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
		visit(cur.Data)
		// 对右子树进行 goAlongLeftBranch 操作
		cur = cur.Right
	}
}

// highest left visible leaf
func gotoHLVL(s *[]*TreeNode) {
	cur := (*s)[len(*s)-1]
	for ; cur != nil; cur = (*s)[len(*s)-1] {
		// 尽量向左子树深入
		if cur.Left != nil {
			// 将这一层的左右子节点都入栈
			if cur.Right != nil {
				*s = append(*s, cur.Right)
			}
			*s = append(*s, cur.Left)
		} else {
			// 实在不行，才向右深入
			*s = append(*s, cur.Right)
		}
	}
	// 去掉栈顶的nil
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
		visit(cur.Data)
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
