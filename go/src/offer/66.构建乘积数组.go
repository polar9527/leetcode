package main

func constructArr(a []int) []int {
	l := len(a)
	if l == 0 {
		return nil
	}

	b := make([]int, l)
	b[0] = 1
	for i := 1; i < l; i++ {
		b[0] *= a[i]
	}
	for i := 1; i < l; i++ {
		if a[i] != 0 {
			b[i] = (b[i-1] * a[i-1]) / a[i]
		} else {
			b[i] = 1
			for j := 0; j < l; j++ {
				if j != i {
					b[i] *= a[j]
				}
			}
		}
	}
	return b
}

// func constructArr(a []int) []int {
// 	if a == nil || len(a) == 0 {
// 		return nil
// 	}
// 	left, right := make([]int, len(a)), make([]int, len(a))
// 	left[0], right[len(a) - 1] = 1, 1
// 	for i := 1; i < len(a); i++ {
// 		left[i] = left[i-1] * a[i-1]
// 	}
// 	for i := len(a) - 2; i >= 0; i-- {
// 		right[i] = right[i + 1] * a[i+1]
// 	}
// 	ret := make([]int, len(a))
// 	for i, _ := range a {
// 		ret[i] = left[i] * right[i]
// 	}
// 	return ret
// }

func main() {
	constructArr([]int{1, 2, 3, 4, 5})
}
