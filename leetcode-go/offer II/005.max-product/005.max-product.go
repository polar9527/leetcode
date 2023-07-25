package offer2

func maxProduct(words []string) int {
	ret := 0
	bitMap := make(map[int]int)
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		if len(word) > bitMap[mask] {
			bitMap[mask] = len(word)
		}
	}

	for xk, xl := range bitMap {
		for yk, yl := range bitMap {
			if xk&yk == 0 && xl*yl > ret {
				ret = xl * yl
			}
		}
	}
	return ret
}
