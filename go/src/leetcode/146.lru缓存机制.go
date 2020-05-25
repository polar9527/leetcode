// /*
//  * @lc app=leetcode.cn id=146 lang=golang
//  *
//  * [146] LRU缓存机制
//  *
//  * https://leetcode-cn.com/problems/lru-cache/description/
//  *
//  * algorithms
//  * Medium (47.28%)
//  * Likes:    615
//  * Dislikes: 0
//  * Total Accepted:    63.1K
//  * Total Submissions: 129.2K
//  * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
//   '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
//  *
//  * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//  *
//  * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//  * 写入数据 put(key, value) -
//  * 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//  *
//  *
//  *
//  * 进阶:
//  *
//  * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
//  *
//  *
//  *
//  * 示例:
//  *
//  * LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//  *
//  * cache.put(1, 1);
//  * cache.put(2, 2);
//  * cache.get(1);       // 返回  1
//  * cache.put(3, 3);    // 该操作会使得密钥 2 作废
//  * cache.get(2);       // 返回 -1 (未找到)
//  * cache.put(4, 4);    // 该操作会使得密钥 1 作废
//  * cache.get(1);       // 返回 -1 (未找到)
//  * cache.get(3);       // 返回  3
//  * cache.get(4);       // 返回  4
//  *
//  *
//  */
package main

// @lc code=start
type LRUCache struct {
	head, tail *DLinkNode
	cache      map[int]*DLinkNode
	capacity   int
	size       int
}

type DLinkNode struct {
	pre        *DLinkNode
	next       *DLinkNode
	key, value int
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*DLinkNode)
	head := initDLinkedNode(0, 0)
	tail := initDLinkedNode(0, 0)
	head.next = tail
	tail.pre = head
	return LRUCache{head, tail, cache, capacity, 0}
}

func initDLinkedNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func (this *LRUCache) Get(key int) int {
	if ptr, ok := this.cache[key]; ok {
		this.moveToHead(ptr)
		return ptr.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ptr, ok := this.cache[key]; ok {
		this.moveToHead(ptr)
		ptr.value = value
	} else {
		this.size++
		newNode := initDLinkedNode(key, value)
		this.cache[key] = newNode
		this.addToHead(newNode)
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	}

}

func (this *LRUCache) moveToHead(node *DLinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *DLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addToHead(node *DLinkNode) {
	l := this.head
	r := this.head.next
	l.next = node
	node.pre = l
	node.next = r
	r.pre = node
}

func (this *LRUCache) removeTail() *DLinkNode {
	removed := this.tail.pre
	this.removeNode(removed)
	return removed
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func main() {

}
