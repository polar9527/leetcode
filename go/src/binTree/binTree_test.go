package bintree

import "testing"

func TestBinTree(t *testing.T) {
	tree := &Tree{}
	tree.Add(50)
	tree.Add(45)
	tree.Add(40)
	tree.Add(48)
	tree.Add(51)
	tree.Add(61)
	tree.Add(71)

	t.Log("preorderTraverse")
	t.Log("Recursion")
	tree.preorderTraverseRecursion(tree.root)
	t.Log()
	t.Log("Iteration")
	tree.preorderTraverseIteration(tree.root)

	t.Log("")

	t.Log("inorderTraverse")
	t.Log("Recursion")
	tree.inorderTraverseRecursion(tree.root)
	t.Log()
	t.Log("Iteration")
	tree.inorderTraverseIteration(tree.root)

	t.Log("")

	t.Log("postorderTraverse")
	t.Log("Recursion")
	tree.postorderTraverseRecursion(tree.root)
	t.Log()
	t.Log("Iteration")
	tree.postorderTraverseIteration(tree.root)

}
