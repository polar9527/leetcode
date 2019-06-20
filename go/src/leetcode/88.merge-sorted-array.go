/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
// package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	p := m + n - 1
	m--
	n--
	for i := p; i >= 0; i-- {
		if m >= 0 && n >= 0 {
			if nums1[m] >= nums2[n] {
				nums1[i] = nums1[m]
				m--
				continue
			} else {
				nums1[i] = nums2[n]
				n--
				continue
			}
		}

		if n < 0 {
			nums1[i] = nums1[m]
			m--
			continue
		}

		if m < 0 {
			nums1[i] = nums2[n]
			n--
			continue
		}

	}
}

// func main() {
// 	nums1 := []int{1, 2, 3, 0, 0, 0}
// 	m := 3
// 	nums2 := []int{2, 5, 6}
// 	n := 3

// 	merge(nums1, m, nums2, n)

// }
