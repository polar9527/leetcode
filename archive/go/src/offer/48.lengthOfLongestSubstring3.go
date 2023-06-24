package main

func lengthOfLongestSubstring(s string) int {
	b := []byte(s)
	set := make(map[byte]int)
	tail, head, max := 0, 0, 0
	for ; head < len(b); head++ {
		if _, ok := set[b[head]]; !ok {
			set[b[head]] = head
		} else {
			if set[b[head]] == tail {
				set[b[head]] = head
				tail++
			} else {
				// remove
				for start := tail; start < set[b[head]]; start++ {
					delete(set, b[start])
				}
				tail = set[b[head]] + 1
				set[b[head]] = head
			}
		}
		l := head - tail + 1
		if max < l {
			max = l
		}
	}
	return max
}

func main() {

	lengthOfLongestSubstring("abcabcbb")
}
