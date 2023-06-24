package main

func movingCount(m int, n int, k int) int {
	if k < 0 || m <= 0 || n <= 0 {
		return 0
	}

	visited := make([]bool, m*n)

	count := movingCountRicursive(k, m, n, 0, 0, &visited)
	return count
}

func movingCountRicursive(threshold, rows, cols, r, c int, visited *[]bool) int {

	count := 0
	if check(threshold, rows, cols, r, c, visited) {
		(*visited)[r*cols+c] = true
		right := movingCountRicursive(threshold, rows, cols, r+1, c, visited)
		left := movingCountRicursive(threshold, rows, cols, r-1, c, visited)
		up := movingCountRicursive(threshold, rows, cols, r, c-1, visited)
		down := movingCountRicursive(threshold, rows, cols, r, c+1, visited)
		count = 1 + right + left + up + down
	}

	return count

}

func check(threshold, rows, cols, r, c int, visited *[]bool) bool {
	if r >= 0 && r < rows && c >= 0 && c < cols && !(*visited)[r*cols+c] && getDigitSum(r)+getDigitSum(c) <= threshold {
		return true
	}
	return false
}

func getDigitSum(num int) int {
	sum := 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func main() {
	movingCount(1, 2, 1)
}
