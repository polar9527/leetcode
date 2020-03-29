package main

func cuttingRopeDynamic(n int) int {
	switch {
	case n < 2:
		return 0
	case n == 2:
		return 1
	case n == 3:
		return 2
	default:
	}

	dynamicMap := make(map[int]int)
	dynamicMap[0] = 0
	dynamicMap[1] = 1
	dynamicMap[2] = 2
	dynamicMap[3] = 3
	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			product := dynamicMap[j] * dynamicMap[i-j]
			if max < product {
				max = product
			}
		}
		dynamicMap[i] = max
	}
	return dynamicMap[n]
}

// func main(){
// 	cuttingRopeDynamic(10)
// }
