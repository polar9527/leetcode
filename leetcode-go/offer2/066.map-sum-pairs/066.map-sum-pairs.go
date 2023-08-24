package offer2

/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 *
 * https://leetcode.cn/problems/map-sum-pairs/description/
 *
 * algorithms
 * Medium (65.61%)
 * Likes:    241
 * Dislikes: 0
 * Total Accepted:    49.2K
 * Total Submissions: 74.9K
 * Testcase Example:  '["MapSum","insert","sum","insert","sum"]\n' +
  '[[],["apple",3],["ap"],["app",2],["ap"]]'
 *
 * 设计一个 map ，满足以下几点:
 *
 *
 * 字符串表示键，整数表示值
 * 返回具有前缀等于给定字符串的键的值的总和
 *
 *
 * 实现一个 MapSum 类：
 *
 *
 * MapSum() 初始化 MapSum 对象
 * void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键
 * key 已经存在，那么原来的键值对 key-value 将被替代成新的键值对。
 * int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 * ["MapSum", "insert", "sum", "insert", "sum"]
 * [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
 * 输出：
 * [null, null, 3, null, 5]
 *
 * 解释：
 * MapSum mapSum = new MapSum();
 * mapSum.insert("apple", 3);
 * mapSum.sum("ap");           // 返回 3 (apple = 3)
 * mapSum.insert("app", 2);
 * mapSum.sum("ap");           // 返回 5 (apple + app = 3 + 2 = 5)
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= key.length, prefix.length <= 50
 * key 和 prefix 仅由小写英文字母组成
 * 1 <= val <= 1000
 * 最多调用 50 次 insert 和 sum
 *
 *
*/

// @lc code=start
// type MapSum map[string]int

// func Constructor() MapSum {
// 	return MapSum{}
// }

// func (m MapSum) Insert(key string, val int) {
// 	m[key] = val
// }

// func (m MapSum) Sum(prefix string) (sum int) {
// 	for s, v := range m {
// 		if strings.HasPrefix(s, prefix) {
// 			sum += v
// 		}
// 	}
// 	return
// }

type TrieNode struct {
	children [26]*TrieNode
	val      int
}

type MapSum struct {
	root *TrieNode
	cnt  map[string]int
}

func Constructor() MapSum {
	return MapSum{&TrieNode{}, map[string]int{}}
}

func (m *MapSum) Insert(key string, val int) {
	delta := val
	if m.cnt[key] > 0 {
		delta -= m.cnt[key]
	}
	m.cnt[key] = val
	node := m.root
	for _, ch := range key {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
		node.val += delta
	}
}

func (m *MapSum) Sum(prefix string) int {
	node := m.root
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return 0
		}
		node = node.children[ch]
	}
	return node.val
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end
