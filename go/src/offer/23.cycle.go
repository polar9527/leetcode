package main

func getCyclePoint(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 	whether has cycle
	var meet *ListNode

	slowStep := head
	fastStep := head.Next
	if fastStep == nil {
		return nil
	}
	for fastStep != nil && slowStep != nil {
		if slowStep == fastStep {
			meet = fastStep
			break
		}
		slowStep = slowStep.Next

		fastStep = fastStep.Next
		if fastStep != nil {
			fastStep = fastStep.Next
		}
	}
	if meet == nil {
		return nil
	}

	// size of cycle
	cycleSize := 1
	node := meet
	for node.Next != meet {
		node = node.Next
		cycleSize++
	}

	// 	entry point
	first := head
	second := head
	for i := 0; i < cycleSize; i++ {
		first = first.Next
	}

	for first != second {
		first = first.Next
		second = second.Next
	}
	return second
}
