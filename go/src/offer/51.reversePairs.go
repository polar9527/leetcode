package main

func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return reversePairsCore(nums)
}

func reversePairsCore(data []int) int {
	if len(data) == 1 {
		return 0
	}
	mid := len(data) / 2
	// divide
	left := reversePairsCore(data[:mid])
	right := reversePairsCore(data[mid:])
	count := 0
	// merger left and right
	start := 0
	end := len(data) - 1
	// reset cache
	cache := make([]int, 0)
	i, j := mid-1, end
	for i >= start && j >= mid {
		if data[i] > data[j] {
			// 	有逆序
			count += j - mid + 1
			cache = append(cache, data[i])
			i--
		} else {
			cache = append(cache, data[j])
			j--
		}
	}
	for i >= start {
		cache = append(cache, data[i])

		i--
	}
	for j >= mid {
		cache = append(cache, data[j])
		j--
	}
	l := len(data) - 1
	for i, n := range cache {
		data[l-i] = n
	}

	return count + left + right
}

func main() {
	nums := []int{7, 5, 6, 4}
	reversePairs(nums)
}
