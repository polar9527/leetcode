package main

func spiralOrder(matrix [][]int) []int {

	var ret []int
	rows := len(matrix)
	if rows == 0 {
		return ret
	}
	cols := len(matrix[0])
	if cols == 0 {
		return ret
	}
	var cycle []int
	start := 0
	for rows > start*2 && cols > start*2 {
		cycle = getCycle(matrix, start)
		ret = append(ret, cycle...)
		start++
	}

	return ret
}

func getCycle(matrix [][]int, start int) []int {

	var ret []int
	rows := len(matrix)
	cols := len(matrix[0])
	rowEnd := rows - 1 - start
	colEnd := cols - 1 - start

	for i := start; i <= colEnd; i++ {
		ret = append(ret, matrix[start][i])
	}

	if rowEnd > start {
		for i := start + 1; i <= rowEnd; i++ {
			ret = append(ret, matrix[i][colEnd])
		}
	}

	if rowEnd > start && colEnd > start {
		for i := colEnd - 1; i >= start; i-- {
			ret = append(ret, matrix[rowEnd][i])
		}
	}
	if rowEnd-1 > start && colEnd > start {
		for i := rowEnd - 1; i >= start+1; i-- {
			ret = append(ret, matrix[i][start])
		}
	}

	return ret
}

func main() {
	m := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	spiralOrder(m)
}
