/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 *
 * https://leetcode-cn.com/problems/word-search-ii/description/
 *
 * algorithms
 * Hard (45.35%)
 * Likes:    400
 * Dislikes: 0
 * Total Accepted:    37K
 * Total Submissions: 81.6K
 * Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n' +
  '["oath","pea","eat","rain"]'
 *
 * 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。
 *
 * 单词必须按照字母顺序，通过 相邻的单元格
 * 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board =
 * [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
 * words = ["oath","pea","eat","rain"]
 * 输出：["eat","oath"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["a","b"],["c","d"]], words = ["abcb"]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n == board[i].length
 * 1
 * board[i][j] 是一个小写英文字母
 * 1
 * 1
 * words[i] 由小写英文字母组成
 * words 中的所有字符串互不相同
 *
 *
*/

// @lc code=start
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func TrieConstructor() *Trie {
	return &Trie{}
}

func (t *Trie) insert(word string) *Trie {
	node := t
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Trie{}
		}
		node = node.children[c-'a']
	}
	node.isEnd = true
	return t
}

func (t *Trie) search(word string) bool {
	node := t
	for _, c := range word {
		if t.children[c-'a'] == nil {
			return false
		}
		node = t.children[c-'a']
	}
	return node.isEnd

}

func findWords(board [][]byte, words []string) []string {
	xLen := len(board)
	if xLen < 0 {
		return nil
	}
	yLen := len(board[0])

	var res []string
	var trace []byte
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	t := TrieConstructor()
	for _, w := range words {
		t.insert(w)
	}

	var dfs func(x, y int, parent *Trie)
	dfs = func(x, y int, parent *Trie) {
		if x < 0 || x >= xLen || y < 0 || y >= yLen || board[x][y] == '.' {
			return
		}

		trace = append(trace, board[x][y])
		char := board[x][y]
		board[x][y] = byte('.')
		defer func() {
			trace = trace[:len(trace)-1]
			board[x][y] = char
		}()
		child := parent.children[char-'a']
		if child == nil {
			return
		}
		if child.isEnd {
			res = append(res, string(trace))
			child.isEnd = false
		}
		for _, d := range direction {
			dfs(x+d[0], y+d[1], child)
		}
	}

	for i := range board {
		for j := range board[i] {
			dfs(i, j, t)
		}
	}
	return res
}

// @lc code=end

