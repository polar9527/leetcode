package main

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	res := []string{}
	permutationCore([]byte(s), 0, &res)
	return res
}

func permutationCore(s []byte, index int, res *[]string) {
	if index == len(s) {
		*res = append(*res, string(s))
		return
	}
	m := make(map[byte]bool)
	for i := index; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			continue
		}
		s[index], s[i] = s[i], s[index]
		permutationCore(s, index+1, res)
		s[index], s[i] = s[i], s[index]
		m[s[i]] = true
	}

}
