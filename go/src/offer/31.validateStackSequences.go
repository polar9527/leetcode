package main

func validateStackSequences(pushed []int, popped []int) bool {

	l := len(pushed)
	if l == 0 && len(popped) != 0 {
		return false
	}
	stack := make([]int, 0, l)
	pushIndex := 0
	for i := 0; i < l; {
		if len(stack) == 0 || stack[len(stack)-1] != popped[i] {
			if pushIndex == l {
				return false
			}
			stack = append(stack, pushed[pushIndex])
			pushIndex++

		} else {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	return true
}
