package acm

var n = 1005
var fatherMap = make(map[int]int)

func init() {
	for i := 0; i < n; i++ {
		fatherMap[i] = i
	}
}

func find(key int) int {
	if key == fatherMap[key] {
		return key
	}
	// 递归查找
	root := find(fatherMap[key])
	// 压缩路径
	fatherMap[key] = root
	return root

}

// Is the same root?
func isSame(u, v int) bool {
	uRoot := find(u)
	vRoot := find(v)
	return uRoot == vRoot
}

func join(son, father int) {
	sonRoot := find(son)
	fatherRoot := find(father)
	if sonRoot == fatherRoot {
		return
	}
	fatherMap[son] = father
}
