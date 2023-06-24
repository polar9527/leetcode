package sort

func mergeSort(arr []int) []int {
	l := len(arr)

	if l < 2 {
		return arr
	}

	middle := l / 2
	merge(mergeSort(arr[:middle]), mergeSort(arr[middle:]), arr)
	return arr
}

func merge(left, right, all []int) {
	temp := make([]int, len(left))
	copy(temp, left)
	a, l, r := 0, 0, 0
	for l < len(left) && r < len(right) {
		if temp[l] < right[r] {
			all[a] = temp[l]
			a++
			l++
		} else {
			all[a] = right[r]
			a++
			r++
		}
	}
	for l < len(left) {
		all[a] = temp[l]
		a++
		l++
	}
	for r < len(right) {
		all[a] = right[r]
		a++
		r++
	}

	return
}

func mergeNotInPlace(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0 // left和right的index位置
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
		} else {
			result = append(result, left[m])
			m++
		}
	}
	if n < r {
		result = append(result, right[n:]...)
	}
	if m < l {
		result = append(result, left[m:]...)
	}
	return result
}
