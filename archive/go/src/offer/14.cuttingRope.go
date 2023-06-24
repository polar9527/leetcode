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

func cuttingRopeGreedy(n int) int {
	switch {
	case n < 2:
		return 0
	case n == 2:
		return 1
	case n == 3:
		return 2
	default:
	}
	p := 1000000007
	// 尽可能多地减去长度为3的绳子段
	timesOf3 := n / 3

	// 当绳子最后剩下的长度为4的时候，不能再剪去长度为3的绳子段。
	// 此时更好的方法是把绳子剪成长度为2的两段，因为2*2 > 3*1。
	if n-timesOf3*3 == 1 {
		timesOf3 -= 1
	}

	timesOf2 := (n - timesOf3*3) / 2

	p3 := pow(3, timesOf3)
	p2 := pow(2, timesOf2)
	ret := (p3 * p2) % p
	return ret

}

func pow(a, x int) int {
	p := 1000000007

	acc := 1
	for i := 0; i < x; i++ {
		acc = (acc * a) % p
	}
	return acc
}

// func main(){
// 	retd := cuttingRopeDynamic(120)
// 	retg := cuttingRopeGreedy(120)
// 	fmt.Println(retd, retg)
// }
