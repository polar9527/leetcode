package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Tuple struct {
	array  []int
	target int
}

type BinSearch func([]int, int) int

func GetFunctionName(i interface{}, seps ...rune) string {
	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer())
	fnName := fn.Name()

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fnName, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	// fmt.Println(fields)

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

// 在一个有序数组arr中, 寻找大于等于target的元素的第一个索引，如果存在, 则返回相应的索引index，否则, 返回arr的元素个数 n。
func main() {
	var testcase []Tuple
	// testcase = append(testcase, Tuple{[]int{1, 3, 6, 13, 14, 14, 14, 14, 56, 86, 99, 111}, 14})
	// testcase = append(testcase, Tuple{[]int{1, 3, 6, 12, 13, 13, 13, 13, 13, 56, 86, 99, 111}, 14})
	// testcase = append(testcase, Tuple{[]int{111}, 14})
	// testcase = append(testcase, Tuple{[]int{11}, 14})
	// testcase = append(testcase, Tuple{[]int{}, 14})
	testcase = append(testcase, Tuple{[]int{1, 2, 5, 6, 7, 8, 9}, 4})

	// testBinSearch(binSearchExactlyHalfClose, testcase)
	// testBinSearch(binSearchExactlyFullClose, testcase)
	// testBinSearch(binSearchLowerBound, testcase)
	// testBinSearch(binSearchUpperBound, testcase)
	testBinSearch(bin, testcase)

}

func testBinSearch(f BinSearch, testcase []Tuple) {

	fName := GetFunctionName(f, '.')
	fmt.Println("Function ====> ", fName)
	for i, t := range testcase {
		fmt.Println("testcase ==> ", i)
		fmt.Print("array ==> ", t.array, "\t", "target ==> ", t.target)
		fmt.Println()
		ret := f(t.array, t.target)
		if ret == len(t.array) {
			fmt.Println("Not found: ", ret)
		} else {
			fmt.Println("Found index: ", ret)
		}
	}
	fmt.Println()
}

// wrong during byteDance interview
func binSearch(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}
	if len(a) == 1 {
		if a[0] == target {
			return a[0]
		}
		return -1
	}
	lo, hi, mi := 0, len(a)-1, len(a)/2
	for ; lo < hi && a[mi] != target; mi = lo + (hi-lo)/2 {
		if a[mi] < target {
			lo = mi
		} else if a[mi] >= target {
			hi = mi + 1
		}
	}
	if a[mi] >= target {
		return a[mi]
	} else {
		return -1
	}
}

func bin(a []int, t int) int {
	l := len(a)
	if l == 0 {
		return -1
	}

	lo, hi := 0, l-1
	for lo <= hi {
		mi := lo + (hi-lo)/2
		if a[mi] == t {
			return mi
		}
		if a[mi] < t {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return -1
}

// [l, r)
// 半开半闭
// 极端情况的临界，即 l+1==r, 进入循环
// 进入循环后，mid的值是(2l+1)/2 => l
// 如果任然没有找到target, 即满足条件 target == arr[mid], 此时所在的区间是[l, l+1), 且mid==l， 数组中最后一个待对比元素arr[mid]也与target不等。
// 则无非两种情况，最终不管是 l=mid+1 或者 r=mid, 都会得到l==r 而跳出循环
func binSearchExactlyHalfClose(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if target == arr[mid] {
			return mid
		}
		// 此时[l, mid]这个区间的元素全部小于target，target只可能存在于区间[mid + 1, r)
		if target > arr[mid] {
			l = mid + 1
			// target < arr[mid]
			// 此时[mid, r)这个区间的元素全部大于target, target只可能存在于区间[l, mid-1],
			// 将全闭合区间[l, mid-1]调整为半开半闭区间[l, mid)与循环条件语义的上下文匹配
		} else {
			r = mid
		}
	}
	return -1
}

// [l, r]
// 全闭
// 极端情况的临界， 即 l==r, 进入循环
// 进入循环后， mid的值是l，同时也是r
// 如果任然没有找到target, 即满足条件 target == arr[mid], 此时所在的区间是[l, l+1), 且mid==l， 数组中最后一个待对比元素arr[mid]也与target不等。
// 则无非两种情况，最终不管是 l=mid+1 或者 r=mid-1, 都会得到l>r 而跳出循环
func binSearchExactlyFullClose(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if target == arr[mid] {
			return mid
		}
		// 此时[l, mid]这个区间的元素全部小于target，target只可能存在于区间[mid + 1, r)
		if target > arr[mid] {
			l = mid + 1
			// target < arr[mid]
			// 此时[mid, r)这个区间的元素全部大于target, target只可能存在于全闭合区间[l, mid-1],
		} else {
			r = mid - 1
		}
	}
	return -1
}

// [l, r]
// 全闭
// 在一个有序的数组arr中，寻找大于等于target的元素的第一个索引，如果存在返回相应索引，否则, 返回arr的元素个数
// 极端情况的临界， 即 l==r, 进入循环
// 此时 mid == l == r
// 如果target == arr[mid]，
// 那么这时可以肯定 l (mid == l == r)的左侧的所有元素全部小于target，
// 而arr[mid]此时就是第一个等于target的元素
// 如果target < arr[mid]，
// 那么这时可以肯定 r (mid == l == r)的右侧的所有元素全部大于target，
// 同时此时 l (mid == l == r)的左侧的所有元素全部小于target
// 则arr[mid]此时就是第一个大于target的元素
// 然后有r=mid-1, l=mid, l>r 循环结束。
//
// 如果arr[mid] < target,
// 那么这时可以肯定 l (mid == l == r) 的左侧的所有元素全部小于target，
// 同时又有 r (mid == l == r)的右侧的所有元素全部大于target，
// 那么arr[r+1](即arr[l+1])是第一个大于target的元素
// 然后有r=mid, l=mid+1, l>r 循环结束。
func binSearchLowerBound(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		// 当arr[mid] == target的时候，此时[mid, r]中的元素一定大于等于target，但是arr[mid]可能不是第一个等于target的元素,还需要继续在[l, mid-1]中寻找
		// 当arr[mid] > target的时候，此时[mid, r] 中的元素一定大于target，但是arr[mid]可能不是第一个大于等于target的元素,还需要继续在[l, mid-1]中寻找
		if target <= arr[mid] {
			r = mid - 1
			// 当arr[mid] < target的时候, [l, mid]中的元素全部小于target，不满足条件，
			// 此时大于等于target的元素只存在与[mid+1, r]区间中，
		} else {
			l = mid + 1
		}
	}
	return l
}

// [l, r]
// 全闭
// 在一个有序数组arr中, 寻找大于target的元素的第一个索引，如果存在, 则返回相应的索引index，否则, 返回arr的元素个数

func binSearchUpperBound(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		// [l, mid]区间都小于等于target, 继续到[mid+1, r]区间寻找
		// 临界 l==mid==r
		// l左侧包括l都小于target, r右侧都大于target， 所以arr[r+1]是第一个大于target的
		if arr[mid] <= target {
			l = mid + 1
			// target < arr[mid]
			// [mid, r]区间都大于target， 继续到[l, mid-1]区间寻找
			// 临界 l==mid==r
			// r右侧包括r都大于target，l左侧都小于target，所以arr[l]是第一个大于target的
		} else {
			r = mid - 1
		}
	}
	// r 是最后一个 target 的位置
	// return r
	return l
}
