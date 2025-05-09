package go_case

import "container/list"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 *
 * https://leetcode.cn/problems/lru-cache/description/
 *
 * algorithms
 * Medium (53.46%)
 * Likes:    2793
 * Dislikes: 0
 * Total Accepted:    506.1K
 * Total Submissions: 946.6K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
 *
 * 实现 LRUCache 类：
 *
 *
 *
 *
 * LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 * void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
 * key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
 *
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * 输出
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * 解释
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // 缓存是 {1=1}
 * lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
 * lRUCache.get(1);    // 返回 1
 * lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
 * lRUCache.get(2);    // 返回 -1 (未找到)
 * lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
 * lRUCache.get(1);    // 返回 -1 (未找到)
 * lRUCache.get(3);    // 返回 3
 * lRUCache.get(4);    // 返回 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= capacity <= 3000
 * 0 <= key <= 10000
 * 0 <= value <= 10^5
 * 最多调用 2 * 10^5 次 get 和 put
 *
 *
*/

// @lc code=start
// type LRUCache struct {
// 	head, tail *DLinkNode
// 	cache      map[int]*DLinkNode
// 	capacity   int
// 	size       int
// }

// type DLinkNode struct {
// 	pre        *DLinkNode
// 	next       *DLinkNode
// 	key, value int
// }

// func Constructor(capacity int) LRUCache {
// 	cache := make(map[int]*DLinkNode)
// 	head := initDLinkedNode(0, 0)
// 	tail := initDLinkedNode(0, 0)
// 	head.next = tail
// 	tail.pre = head
// 	return LRUCache{head, tail, cache, capacity, 0}
// }

// func initDLinkedNode(key, value int) *DLinkNode {
// 	return &DLinkNode{
// 		key:   key,
// 		value: value,
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	if ptr, ok := this.cache[key]; ok {
// 		this.moveToHead(ptr)
// 		return ptr.value
// 	}
// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	if ptr, ok := this.cache[key]; ok {
// 		this.moveToHead(ptr)
// 		ptr.value = value
// 	} else {
// 		this.size++
// 		newNode := initDLinkedNode(key, value)
// 		this.cache[key] = newNode
// 		this.addToHead(newNode)
// 		if this.size > this.capacity {
// 			removed := this.removeTail()
// 			delete(this.cache, removed.key)
// 			this.size--
// 		}
// 	}

// }

// func (this *LRUCache) moveToHead(node *DLinkNode) {
// 	this.removeNode(node)
// 	this.addToHead(node)
// }

// func (this *LRUCache) removeNode(node *DLinkNode) {
// 	node.pre.next = node.next
// 	node.next.pre = node.pre
// }

// func (this *LRUCache) addToHead(node *DLinkNode) {
// 	l := this.head
// 	r := this.head.next
// 	l.next = node
// 	node.pre = l
// 	node.next = r
// 	r.pre = node
// }

// func (this *LRUCache) removeTail() *DLinkNode {
// 	removed := this.tail.pre
// 	this.removeNode(removed)
// 	return removed
// }

type LRUCache struct {
	capacity  int
	list      *list.List
	keyToNode map[int]*list.Element
}

type item struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		list.New(),
		map[int]*list.Element{},
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.keyToNode[key]
	if node == nil {
		return -1
	}
	this.list.MoveToFront(node)
	return node.Value.(item).value
}

func (this *LRUCache) Put(key int, value int) {
	node := this.keyToNode[key]
	if node != nil {
		// node.prev = dummy
		// node.next = dummy.Next
		// node.prev.next = node
		// node.next.prev = node
		this.list.MoveToFront(node)
		node.Value = item{key, value}
		return
	}
	node = this.list.PushFront(item{key, value})
	this.keyToNode[key] = node
	if len(this.keyToNode) > this.capacity {
		deletedNode := this.list.Back()
		di := this.list.Remove(deletedNode).(item)
		delete(this.keyToNode, di.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
