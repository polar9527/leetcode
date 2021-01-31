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

func insertionSort1(arr []int) {
	l := len(arr)

	if l < 2 {
		return
	}

	for i := 1; i < l; i++ {
		key := arr[i]
		// A[0, i-1] sorted
		// A[i:] unsorted
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
