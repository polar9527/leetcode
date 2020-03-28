package main

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]
	}

	lo := 0
	hi := len(numbers) - 1
	mi := lo
	for numbers[lo] >= numbers[hi] {
		if lo+1 == hi {
			mi = hi
			break
		}
		mi = (lo + hi) / 2
		if numbers[mi] == numbers[hi] && numbers[mi] == numbers[lo] {
			min := numbers[lo]
			for i := lo; i <= hi; i++ {
				if numbers[i] < min {
					min = numbers[i]
				}
			}
			return min
		}
		if numbers[mi] <= numbers[hi] {
			hi = mi
		} else if numbers[mi] >= numbers[lo] {
			lo = mi
		}
	}
	return numbers[mi]
}

// func main() {
// 	a := []int{3,3,1,3}
// 	ret := minArray(a)
// 	fmt.Println(ret)
// }
