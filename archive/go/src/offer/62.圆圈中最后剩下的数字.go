package main

// n=8	m=3
//
// n=8		0  1  2  3  4  5  6  7		g 在当前序列的序号
// 		    a  b [c] d  e  f {g} h		F(8, 3) = 6
// n=7		0  1  2  3  4  5  6  7
// 		    d  e [f]{g} h  a  b			F(7, 3) = 3
// n=6		0  1  2  3  4  5  6  7
//         {g} h [a] b  d  e			F(6, 3) = 0
// n=5		0  1  2  3  4  5  6  7
// 		    b  d [e]{g} h				F(5, 3) = 3
// n=4		0  1  2  3  4  5  6  7
// 	       {g}  h [b] d					F(4, 3) = 0
// n=3		0  1  2  3  4  5  6  7
// 		    d {g}[h]					F(3, 3) = 1
// n=2		0  1  2  3  4  5  6  7
//         [d]{g}						F(2, 3) = 1
// n=1		0  1  2  3  4  5  6  7
// 	       {g}							F(1, 3) = 0
//
// n=6		0  1  2  3  4  5  6  7  8  9 10
//          g  h  a  b  d  e (f)				序列补上上次删除的字符
// 	                 g  h  a  b  d  e (f)	  	每个字符的序号向右移动m=3,    F(6, 3) + m
// 	        d  e  f  g  h  a  b					每个字符的序号模n=7,   	   (F(6, 3) + m) % 7 ==  F(7, 3)
//
// n=7      0  1  2  3  4  5  6  7  8  9 10
// 		    d  e  f  g  h  a  b (c)			    序列补上上次删除的字符
// 		             d  e  f  g  h  a  b (c) 	每个字符的序号向右移动m=3,    F(7, 3) + m
// 		    a  b (c) d  e  f  g  h				每个字符的序号模n=8          (F(7, 3) + m) % 8 ==  F(8, 3)
//
// n=8	    0  1  2  3  4  5  6  7
// 		    a  b [c] d  e  f  g  h
//
// F(n, m) = (F(n-1, m) + m)%n
// n递减, m不变
// F(1, m) == 0
// F(2, m) == (0+m)%2
// F(i, m) == (F(i-1, m) + m)%i

func lastRemaining(n int, m int) int {
	index := 0
	for i := 2; i <= n; i++ {
		index = (index + m) % i
	}
	return index
}
