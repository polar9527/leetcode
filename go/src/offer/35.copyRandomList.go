package main

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	curNode := head
	// insert
	for curNode != nil {
		nextNode := curNode.Next
		newNode := new(Node)
		newNode.Val = curNode.Val

		curNode.Next = newNode
		newNode.Next = nextNode

		curNode = nextNode
	}

	// follow
	curNode = head
	for curNode != nil {
		curNodeClone := curNode.Next
		if curNode.Random != nil {
			curNodeClone.Random = curNode.Random.Next
		}
		curNode = curNodeClone.Next
	}

	// detach
	curNode = head
	cloneHead := head.Next

	curNodeClone := curNode.Next

	for curNode != nil {
		curNode.Next = curNodeClone.Next
		curNode = curNodeClone.Next
		if curNode != nil {
			curNodeClone.Next = curNode.Next
			curNodeClone = curNode.Next
		}
	}
	return cloneHead
}
