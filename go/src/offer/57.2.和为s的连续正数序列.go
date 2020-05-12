package main

func findContinuousSequence(target int) [][]int {
	results := make([][]int, 0)
	end := target/2 + 1

	for tail, head := 1, 1; head <= end; {
		sum := (tail + head) * (head - tail + 1) / 2
		if sum == target {
			result := make([]int, 0)
			for i := tail; i <= head; i++ {
				result = append(result, i)
			}
			results = append(results, result)
			head++
		}
		if sum < target {
			head++
			continue
		}
		if sum > target {
			tail++
		}
	}
	return results
}
