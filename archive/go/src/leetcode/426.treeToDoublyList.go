package leetcode

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var tailNode *TreeNode
	convert(root, &tailNode)

	headNode := root
	for headNode != nil && headNode.Left != nil {
		headNode = headNode.Left
	}
	if tailNode != nil {
		tailNode.Right = headNode
	}
	if headNode != nil {
		headNode.Left = tailNode
	}

	return headNode
}

func convert(root *TreeNode, linkNode **TreeNode) {
	if root == nil {
		return
	}
	currentNode := root
	if currentNode.Left != nil {
		convert(root.Left, linkNode)
	}

	currentNode.Left = *linkNode
	if *linkNode != nil {
		(*linkNode).Right = currentNode
	}

	*linkNode = currentNode

	if currentNode.Right != nil {
		convert(root.Right, linkNode)
	}
}

// func main(){
// 	test1()
// 	test2()
// }
//
// func test1(){
// 	Node4 := &TreeNode{4,nil,nil}
// 	Node2 := &TreeNode{2,nil,nil}
// 	Node5 := &TreeNode{5,nil,nil}
// 	Node4.Left = Node2
// 	Node4.Right = Node5
// 	Node1 := &TreeNode{1,nil,nil}
// 	Node3 := &TreeNode{3,nil,nil}
// 	Node2.Left = Node1
// 	Node2.Right = Node3
//
// 	head := treeToDoublyList(Node4)
//
// 	if head == nil || head.Val != Node1.Val {
// 		fmt.Println("test1 False!")
// 		return
// 	}
//
// 	cur := head
// 	for cur != nil {
// 		fmt.Println(cur.Val)
// 		cur = cur.Right
// 		if cur == head {
// 			fmt.Println(cur.Val)
// 			break
// 		}
// 	}
// }
//
// func test2(){
// 	Node4 := &TreeNode{4,nil,nil}
//
// 	head := treeToDoublyList(Node4)
//
// 	if head == nil || head.Val != Node4.Val {
// 		fmt.Println("test2 False!")
// 		return
// 	}
//
// 	cur := head
// 	for cur != nil {
// 		fmt.Println(cur.Val)
// 		cur = cur.Right
// 		if cur == head {
// 			fmt.Println(cur.Val)
// 			break
// 		}
// 	}
// }
