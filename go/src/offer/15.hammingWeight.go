package main

func hammingWeight(num uint32) int {
	count := 0
	// flag := uint32(1)
	// for flag !=0 {
	// 	if num & flag != 0 {
	// 		count++
	// 	}
	// 	flag = flag << 1
	// }

	for num != 0 {
		count++
		num = (num - 1) & num
	}

	return count
}
