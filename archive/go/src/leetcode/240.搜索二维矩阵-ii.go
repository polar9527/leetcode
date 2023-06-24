/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (46.75%)
 * Likes:    676
 * Dislikes: 0
 * Total Accepted:    148.7K
 * Total Submissions: 317.1K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n' +
  '5'
 *
 * 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
 * 
 * 
 * 每行的元素从左到右升序排列。
 * 每列的元素从上到下升序排列。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：matrix =
 * [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
 * target = 5
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：matrix =
 * [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
 * target = 20
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == matrix.length
 * n == matrix[i].length
 * 1 
 * -10^9 
 * 每行的所有元素从左到右升序排列
 * 每列的所有元素从上到下升序排列
 * -10^9 
 * 
 * 
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
    
    rows := len(matrix)
    if rows == 0 {
        return false
    }
    columns := len(matrix[0])

    if rows == 1 && columns == 1 {
        return matrix[0][0] == target
    }

    if rows == 1 {
        
        l, h := 0, columns-1
        for l <= h {
            mid := (l+h)/2
            if matrix[0][mid] == target {
                return true
            } else if matrix[0][mid] < target {
                l = mid + 1
            } else {
                h = mid - 1
            }
        }
        return false
    }

    if columns == 1 {
        l, h := 0, rows-1
        for l <= h {
            mid := (l+h)/2
            if matrix[mid][0] == target {
                return true
            } else if matrix[mid][0] < target {
                l = mid + 1
            } else {
                h = mid - 1
            }
        }
        return false
    }

    for r, c := rows - 1, 0; r >= 0 && c < columns; {
        if matrix[r][c] == target {
            return true
        } else if target > matrix[r][c] {
            c++
            continue
        } else {
            r--
            continue
        }
    } 

    return false
}
// @lc code=end

