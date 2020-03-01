package sort

func insertionSort(arr []int) {
	l := len(arr)

	if l < 2 {
		return
	}

	// i = precursor length
	// j = start of succeed
	for i := 0; i < l-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}

}
