/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 *
 * https://leetcode.com/problems/copy-list-with-random-pointer/description/
 *
 * algorithms
 * Medium (34.00%)
 * Likes:    2798
 * Dislikes: 623
 * Total Accepted:    376.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[7,null],[13,0],[11,4],[10,2],[1,0]]'
 *
 * A linked list is given such that each node contains an additional random
 * pointer which could point to any node in the list or null.
 *
 * Return a deep copy of the list.
 *
 * The Linked List is represented in the input/output as a list of n nodes.
 * Each node is represented as a pair of [val, random_index] where:
 *
 *
 * val: an integer representing Node.val
 * random_index: the index of the node (range from 0 to n-1) where random
 * pointer points to, or null if it does not point to any node.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
 * Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [[1,1],[2,1]]
 * Output: [[1,1],[2,1]]
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input: head = [[3,null],[3,0],[3,null]]
 * Output: [[3,null],[3,0],[3,null]]
 *
 *
 * Example 4:
 *
 *
 * Input: head = []
 * Output: []
 * Explanation: Given linked list is empty (null pointer), so return null.
 *
 *
 *
 * Constraints:
 *
 *
 * -10000 <= Node.val <= 10000
 * Node.random is null or pointing to a node in the linked list.
 * Number of Nodes will not exceed 1000.
 *
 *
 */
package main

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

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

// @lc code=end
