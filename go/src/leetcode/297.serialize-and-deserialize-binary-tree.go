/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 *
 * https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (45.93%)
 * Likes:    2653
 * Dislikes: 131
 * Total Accepted:    293.8K
 * Total Submissions: 639.3K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 * Serialization is the process of converting a data structure or object into a
 * sequence of bits so that it can be stored in a file or memory buffer, or
 * transmitted across a network connection link to be reconstructed later in
 * the same or another computer environment.
 *
 * Design an algorithm to serialize and deserialize a binary tree. There is no
 * restriction on how your serialization/deserialization algorithm should work.
 * You just need to ensure that a binary tree can be serialized to a string and
 * this string can be deserialized to the original tree structure.
 *
 * Example:
 *
 *
 * You may serialize the following tree:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠    / \
 * ⁠   4   5
 *
 * as "[1,2,3,null,null,4,5]"
 *
 *
 * Clarification: The above format is the same as how LeetCode serializes a
 * binary tree. You do not necessarily need to follow this format, so please be
 * creative and come up with different approaches yourself.
 *
 * Note: Do not use class member/global/static variables to store states. Your
 * serialize and deserialize algorithms should be stateless.
 *
 */
package leetcode

import (
	"math/rand"
	"strconv"
	"strings"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	res  []string
	Flag int
}

func Constructor() Codec {
	flag := rand.Intn(2)
	c := Codec{[]string{}, flag}

	return c
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	switch this.Flag {
	case 0:
		return this.bfsSerialize(root)
	case 1:
		return this.dfsSerialize(root)
	default:
		return ""
	}
}

func (this *Codec) bfsSerialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res []string
	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		curNode := queue[0]
		if curNode != nil {
			res = append(res, strconv.Itoa(curNode.Val))
			queue = append(queue, curNode.Left, curNode.Right)
		} else {
			res = append(res, "null")
		}
		queue = queue[1:]
	}
	return strings.Join(res, ",")

}

func (this *Codec) dfsSerialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + "," + this.dfsSerialize(root.Left) + "," + this.dfsSerialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	switch this.Flag {
	case 0:
		return this.bfsDeserialize(data)
	case 1:
		return this.dfsDeserialize(data)
	default:
		return nil
	}
}

func (this *Codec) bfsDeserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	res := strings.Split(data, ",")
	v, _ := strconv.Atoi(res[0])
	root := &TreeNode{v, nil, nil}
	res = res[1:]
	var queue = []*TreeNode{root}
	for len(queue) > 0 && len(res) > 1 {
		if res[0] != "null" {
			v, _ := strconv.Atoi(res[0])
			queue[0].Left = &TreeNode{v, nil, nil}
			queue = append(queue, queue[0].Left)
		}

		if res[1] != "null" {
			v, _ := strconv.Atoi(res[1])
			queue[0].Right = &TreeNode{v, nil, nil}
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
		res = res[2:]
	}

	return root
}

func (this *Codec) dfsDeserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.res = strings.Split(data, ",")
	return this.dfsDeserializeCore()
}
func (this *Codec) dfsDeserializeCore() *TreeNode {
	node := this.res[0]
	this.res = this.res[1:]
	if node == "null" {
		return nil
	}
	v, _ := strconv.Atoi(node)
	root := &TreeNode{
		Val:   v,
		Left:  this.dfsDeserializeCore(),
		Right: this.dfsDeserializeCore(),
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

// func main(){
// 	node1 := &TreeNode{1,nil,nil}
// 	node2 := &TreeNode{2,nil,nil}
// 	node3 := &TreeNode{3,nil,nil}
// 	node1.Left = node2
// 	node1.Right = node3
// 	node4 := &TreeNode{4,nil,nil}
// 	node5 := &TreeNode{5,nil,nil}
// 	node3.Left = node4
// 	node3.Right = node5
//
// 	flag := 1
// 	cs := Constructor()
// 	cs.Flag = flag
// 	s := cs.serialize(node1)
// 	fmt.Println(s)
// 	d := cs.deserialize(s)
// 	fmt.Println(d)
// }
