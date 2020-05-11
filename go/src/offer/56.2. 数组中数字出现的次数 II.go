package main

// func singleNumber(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
//
// 	bits := make([]int, 32)
// 	for _, num := range nums {
// 		mask := 1
// 		for j := 0; j < 32; j++ {
// 			if num&mask != 0 {
// 				bits[j] += 1
// 			}
// 			mask <<= 1
// 		}
// 	}
// 	number := 0
// 	for i := 31; i >= 0; i-- {
//
// 		if bits[i]%3 != 0 {
// 			number += 1
// 		}
// 		if i == 0 {
// 			break
// 		}
// 		number <<= 1
// 	}
// 	return number
// }

// 数电 卡诺图 输入 输出 状态机
func singleNumber(nums []int) int {
	a := 0
	b := 0
	for _, c := range nums {
		a, b = a&^c|b&c, b&^c|^a&^b&c
	}
	return ^a & b
}

//
// 作者：ilwwli
// 链接：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/solution/zhuang-tai-ji-yu-qia-nuo-tu-wei-yun-suan-de-tui-da/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	singleNumber([]int{3, 4, 3, 3})
}
