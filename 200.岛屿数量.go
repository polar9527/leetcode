/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (49.54%)
 * Likes:    600
 * Dislikes: 0
 * Total Accepted:    115.9K
 * Total Submissions: 233.9K
 * Testcase Example:  '[['1','1','1','1','0'],['1','1','0','1','0'],['1','1','0','0','0'],['0','0','0','0','0']]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 *
 *
 * 示例 1:
 *
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 * 输出: 3
 * 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 *
 *
 */
package main

// @lc code=start
func numIslands(grid [][]byte) int {

}

func numIslandsDisjoinSet(grid [][]byte) int {

}

func numIslandsBFS(grid [][]byte) int {

	var count int
	row := len(grid)
	if row == 0 {
		return count
	}
	var col int
	col = len(grid[0])
	if col == 0 {
		return count
	}

	type pair struct {
		r int
		c int
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' {
				count++
				grid[r][c] = '0'
				dequeu := make([]*pair, 1)
				dequeu[0] = &pair{r, c}
				for len(dequeu) != 0 {
					node := dequeu[0]
					dequeu = dequeu[1:]
					r := node.r
					c := node.c
					list := []*pair{
						&pair{r - 1, c},
						&pair{r + 1, c},
						&pair{r, c - 1},
						&pair{r, c + 1},
					}
					for _, p := range list {
						if p.r >= 0 && p.r < row && p.c >= 0 && p.c < col && grid[p.r][p.c] == '1' {
							grid[p.r][p.c] = '0'
							dequeu = append(dequeu, p)
						}
					}
				}
			}
		}
	}

	return count
}

func numIslandsDFS(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	var col int
	col = len(grid[0])
	if col == 0 {
		return 0
	}

	state := make([]bool, row*col)

	var visit func(int, int) bool
	visit = func(r, c int) bool {
		if r < 0 || r >= row {
			return false
		}
		if c < 0 || c >= col {
			return false
		}
		if grid[r][c] == '0' || state[r*col+c] == true {
			return false
		} else {
			state[r*col+c] = true
			visit(r, c+1)
			visit(r+1, c)
			visit(r-1, c)
			visit(r, c-1)
		}
		return true
	}
	count := 0

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' && state[r*col+c] == false {
				if visit(r, c) {
					count++
				}
			}
		}
	}
	return count
}

// @lc code=end

func main() {
	g := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	// g := [][]byte{
	// 	{'1','1','0','0','0'},
	// 	{'1','1','0','0','0'},
	// 	{'0','0','1','0','0'},
	// 	{'0','0','0','1','1'},
	// }
	// g := [][]byte{
	// 	{'1','0','1'},
	// 	{'0','1','0'},
	// 	{'1','0','1'},
	// }
	numIslands(g)
}
