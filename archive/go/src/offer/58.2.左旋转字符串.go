package main

func reverseLeftWords(s string, n int) string {
	if len(s) == 0 || len(s) < n {
		return ""
	}
	stringByte := []byte(s)
	reverse(stringByte[:n])
	reverse(stringByte[n:])
	reverse(stringByte)

	ret := string(stringByte)
	return ret
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	reverseLeftWords("lrloseumgh", 6)
}
