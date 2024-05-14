package slice_print

import "fmt"

func slicePrint() {
	l := []int{1, 2}
	l = append(l, 3, 4, 5)
	println(cap(l))

	s := make([]int, 5)
	s = append(s, 1, 2, 3, 4, 5)

	s1 := s[2:5]
	s2 := s1[2:6:7]
	fmt.Println(cap(s))
	fmt.Println(cap(s1))
	fmt.Println(cap(s2))

	fmt.Println(s1)
	fmt.Println(s2)

	s3 := append(s2, 6)
	s3[0] = 3

	fmt.Println(s2[0])
	fmt.Println(s)
	fmt.Println(cap(s2))
	fmt.Println(cap(s3))

	s3 = append(s3, 7, 8)
	fmt.Println(s)

}

// 6
// 10
// 8
// 5
// [0 0 0]
// [0 1 2 3]
// 3
// [0 0 0 0 3 1 2 3 6 5]
// 5
// 5
// [0 0 0 0 3 1 2 3 6 5]
