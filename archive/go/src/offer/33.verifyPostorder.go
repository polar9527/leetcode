package main

func verifyPostorder(postorder []int) bool {
	l := len(postorder)
	// 递归基
	if l <= 1 {
		return true
	}
	currentRoot := postorder[l-1]
	rightLeafhead := l - 1
	// 找到左叶子的index
	for i, v := range postorder[:l-1] {
		if v > currentRoot {
			// rightLeafhead >= 0 && rightLeafhead <= l-2
			rightLeafhead = i
			break
		}
	}

	if rightLeafhead < l-1 {
		for _, v := range postorder[rightLeafhead : l-1] {
			if v < currentRoot {
				return false
			}
		}
	}

	return verifyPostorder(postorder[:rightLeafhead]) && verifyPostorder(postorder[rightLeafhead:l-1])

}
