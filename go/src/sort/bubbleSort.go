package sort

func bubbleSort(arr []int) {
	l := len(arr)

	if l < 2 {
		return
	}

	var swapped bool
	for i := 0; i < l-1; i++ {
		swapped = false
		for j := l - 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
