package main

func firstUniqChar(s string) byte {
	result := " "
	bytes := []byte(s)
	dict := make([]int, 256)

	for _, b := range bytes {
		if dict[b] == 0 {
			dict[b] = 1
		} else {
			dict[b] += 1
		}
	}
	for _, b := range bytes {
		if dict[b] == 1 {
			return b
		}
	}
	return result[0]
}
