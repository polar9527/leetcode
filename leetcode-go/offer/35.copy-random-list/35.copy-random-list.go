package offer

import type35 "github.com/polar9527/leetcode/leetcode-go/offer/35.copy-random-list/type"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *type35.Node) *type35.Node {
	if head == nil {
		return nil
	}
	var next *type35.Node
	// round 1: init ego Node for each node in the list
	cur := head
	for cur != nil {
		next = cur.Next
		egoNode := type35.Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next = &egoNode
		cur = next
	}
	// round 2: dump random refers for each node
	cur = head
	for cur != nil {
		if cur.Next == nil {
			panic("what's the fuck here!")
		}
		next = cur.Next.Next
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		} else {
			cur.Next.Random = nil
		}

		cur = next
	}
	// round 3: decouple the ego Node
	cur = head
	ret := head.Next
	for cur != nil {
		if cur.Next == nil {
			panic("what's the fuck here!")
		}
		next = cur.Next.Next
		ego := cur.Next
		cur.Next = ego.Next

		if next != nil {
			ego.Next = next.Next
		}

		cur = next
	}
	return ret
}

/*
   思路：由于Random是随机的，所以每次复制出newNode后，就递归创建newNode.Random，然后递归创建newNode.Next。可以用一个map[*Node]*Node来保存旧链表和新链表每个对应节点
*/
// var record map[*Node]*Node
// func copyRandomList(head *Node) *Node {
//     record = map[*Node]*Node{}

//     return copy(head)
// }

// func copy(node *Node) *Node {
//     if node == nil {
//         return nil
//     }

//     // 如果已经存在，直接返回
//     v, ok := record[node]
//     if ok {
//         return v
//     }

//     // 创建新节点
//     newNode := &Node{
//         Val: node.Val,
//     }

//     // 记录
//     record[node] = newNode

//     // 递归创建random
//     newNode.Random = copy(node.Random)

//     // 创建next
//     newNode.Next = copy(node.Next)

//     return newNode
// }
