package offer2

func twoSum(numbers []int, target int) []int {
    low, high := 0, len(numbers) - 1
    for low < high {
        sum := numbers[low] + numbers[high]
        if sum == target {
            return []int{low, high}
        } else if sum < target {
            low++
        } else {
            high--
        }
    }
    return []int{-1, -1}
}