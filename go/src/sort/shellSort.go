package sort

func shellSort(array []int) {
	l := len(array)
	if l < 2 {
		return
	}
	gap := 1
	for gap < l/3 {
		gap = 3*gap + 1
	}

	for gap > 0 {
		// k = Index In Gap
		for k := 0; k < gap; k++ {
			// j = Index Of Each Gap
			for j := 1; j*gap+k < l; j++ {
				for n := j; n > 0; n-- {
					pre := (n-1)*gap + k
					next := n*gap + k
					if array[pre] > array[next] {
						array[pre], array[next] = array[next], array[pre]
					}
				}
			}
		}
		gap /= 3
	}
}
