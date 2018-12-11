package app

// AppM my solution
func AppM(nums []int) int {

	lenght := len(nums)
	if lenght == 0 || lenght == 1 {
		return lenght
	}
	slower := 0
	for faster := 1; faster < lenght; faster++ {
		if nums[faster] != nums[slower] {
			slower++
			nums[slower] = nums[faster]
		}
	}
	slower++
	nums = nums[:slower]
	return slower
}

//AppO1 a faster solution? Not sure
func AppO1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	nonDup := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nonDup++
			nums[nonDup] = nums[i]
		}
	}
	uniqueNum := nonDup + 1
	nums = nums[:uniqueNum]
	return uniqueNum
}

//AppO2 a faster solution? Not
func AppO2(nums []int) int {
	pmap := make(map[int]int)
	for _, v := range nums {
		if _, ok := pmap[v]; !ok {
			pmap[v] = len(pmap)
			nums[pmap[v]] = v
		}
	}
	return len(pmap)
}
