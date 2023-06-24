/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

package leetcode

func isValidSudoku(board [][]byte) bool {
	boxs := [3][3][]byte{}
	rows := [9][]byte{}
	columns := [9][]byte{}

	dot := []byte(".")[0]

	for iRow, rowCells := range board {
		for iColumn, cell := range rowCells {
			if cell != dot {
				// same row
				for _, vInRow := range rows[iRow] {
					if vInRow == cell {
						return false
					}
				}
				rows[iRow] = append(rows[iRow], cell)
				// same column
				for _, vInColumn := range columns[iColumn] {
					if vInColumn == cell {
						return false
					}
				}
				columns[iColumn] = append(columns[iColumn], cell)
				// same box
				x := iRow / 3
				y := iColumn / 3
				for _, vInBox := range boxs[x][y] {
					if vInBox == cell {
						return false
					}
				}
				boxs[x][y] = append(boxs[x][y], cell)
			}
		}
	}

	return true
}

// func main() {

// 	s := [][]byte{
// 		[]byte("53..7...."),
// 		[]byte("6..195..."),
// 		[]byte(".98....6."),
// 		[]byte("8...6...3"),
// 		[]byte("4..8.3..1"),
// 		[]byte("7...2...6"),
// 		[]byte(".6....28."),
// 		[]byte("...419..5"),
// 		[]byte("....8..79"),
// 	}
// 	print(isValidSudoku(s))
// }
