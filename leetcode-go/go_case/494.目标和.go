package go_case

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 *
 * https://leetcode.cn/problems/target-sum/description/
 *
 * algorithms
 * Medium (48.60%)
 * Likes:    2110
 * Dislikes: 0
 * Total Accepted:    542.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * 给你一个非负整数数组 nums 和一个整数 target 。
 *
 * 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
 *
 *
 * 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
 *
 *
 * 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1,1,1], target = 3
 * 输出：5
 * 解释：一共有 5 种方法让最终目标和为 3 。
 * -1 + 1 + 1 + 1 + 1 = 3
 * +1 - 1 + 1 + 1 + 1 = 3
 * +1 + 1 - 1 + 1 + 1 = 3
 * +1 + 1 + 1 - 1 + 1 = 3
 * +1 + 1 + 1 + 1 - 1 = 3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], target = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 20
 * 0 <= nums[i] <= 1000
 * 0 <= sum(nums[i]) <= 1000
 * -1000 <= target <= 1000
 *
 *
 */

// @lc code=start
// 1D array
// func findTargetSumWays(nums []int, target int) int {
// 	// a + b = sum
// 	// sum - b = a
// 	// a - b = target
// 	// sum - 2b = target
// 	// (sum - target)/2 = b
// 	abs := func(x int) int {
// 		return int(math.Abs(float64(x)))
// 	}

// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	if abs(target) > sum {
// 		return 0
// 	}
// 	if (sum-target)%2 == 1 {
// 		return 0
// 	}
// 	// 计算背包大小
// 	bag := (sum - target) / 2
// 	// 定义dp数组
// 	dp := make([]int, bag+1)
// 	// 初始化
// 	// 即装满背包为0的方法有一种，放0件物品。
// 	dp[0] = 1
// 	// 遍历顺序
// 	for i := 0; i < len(nums); i++ {
// 		for j := bag; j >= nums[i]; j-- {
// 			//推导公式
// 			dp[j] += dp[j-nums[i]]
// 			//fmt.Println(dp)
// 		}
// 	}
// 	return dp[bag]
// }

// 2D array
// func findTargetSumWays(nums []int, target int) int {
// 	abs := func(t int) int {
// 		return int(math.Abs(float64(t)))
// 	}

// 	sum := 0
// 	for _, n := range nums {
// 		sum += n
// 	}
// 	if abs(target) > sum {
// 		return 0
// 	}

// 	if (sum-target)%2 == 1 {
// 		return 0
// 	}
// 	bagSize := (sum - target) / 2

// 	dp := make([][]int, len(nums))
// 	for i, _ := range dp {
// 		dp[i] = make([]int, bagSize+1)
// 	}

// 	// init
// 	dp[0][0] = 1
// 	// first row
// 	if nums[0] <= bagSize {
// 		dp[0][nums[0]] = 1
// 	}
// 	// first column
// 	zeroNums := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			zeroNums++
// 		}
// 		dp[i][0] = int(math.Pow(2.0, float64(zeroNums)))
// 	}

// 	for i := 1; i < len(nums); i++ {
// 		for j := 0; j <= bagSize; j++ {
// 			if j-nums[i] < 0 {
// 				dp[i][j] = dp[i-1][j]
// 			} else {
// 				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
// 			}
// 		}
// 	}
// 	return dp[len(nums)-1][bagSize]
// }

// dfs + memory
// func findTargetSumWays(nums []int, target int) int {
// 	ln := len(nums)
// 	// 所有+号数之和为 plus
// 	// 说有-号数之和为 minus
// 	// 数组全部数字之和为 sum
// 	// sum = plus + minus
// 	// target = plus - minus
// 	// sum + target = 2*plus
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	// plus 是个整数，sum + target = 2*plus
// 	if (sum+target)%2 == 1 {
// 		return 0
// 	}
// 	// 不论 target 是 正数还是负数
// 	// sum + target = 2*plus 都是正数
// 	// 因为 nums 中都是 非负整数， plus 是 nums 中部分 非负整数的和
// 	if target+sum < 0 {
// 		return 0
// 	}

// 	capacity := (sum + target) / 2

// 	cache := make([][]int, ln)
// 	for i := range cache {
// 		// 容量 capacity 的取值范围是 [0, capacity], 一共 capacit+1 个数字
// 		cache[i] = make([]int, capacity+1)
// 		for j := range cache[i] {
// 			cache[i][j] = -1
// 		}
// 	}

// 	// 背包为 plus, 物品为 nums 中的数字，
// 	var dfs func(int, int) int
// 	dfs = func(i, c int) (res int) {
// 		if i < 0 {
// 			if c == 0 {
// 				// 这个时候找到了一个合法的方案
// 				return 1
// 			}
// 			// i 是指 nums中的序号
// 			// 当没有物品可选择时候，即 i < 0
// 			// 没有方案数可以构造 target
// 			return 0
// 		}

// 		// 指针引用 cache 中的结果
// 		p := &cache[i][c]
// 		if *p != -1 {
// 			return *p
// 		}
// 		// 缓存结果
// 		defer func() {
// 			*p = res
// 		}()
// 		if c < nums[i] {
// 			// 无法选择第 i 个数字放入背包
// 			return dfs(i-1, c)
// 		}
// 		// 选择或者不选择第 i 个数字放入背包
// 		// 然后 缩减成 子问题
// 		return dfs(i-1, c) + dfs(i-1, c-nums[i])
// 	}

// 	return dfs(ln-1, capacity)
// }

// DP
// func findTargetSumWays(nums []int, target int) int {
// 	abs := func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}
// 		return x
// 	}
// 	s := 0
// 	for _, x := range nums {
// 		s += x
// 	}
// 	s -= abs(target)
// 	if s < 0 || s%2 == 1 {
// 		return 0
// 	}
// 	m := s / 2 // 背包容量

// 	n := len(nums)
// 	f := make([][]int, n+1)
// 	for i := range f {
// 		f[i] = make([]int, m+1)
// 	}
// 	f[0][0] = 1
// 	for i, x := range nums {
// 		for c := 0; c <= m; c++ {
// 			if c < x {
// 				f[i+1][c] = f[i][c] // 只能不选
// 			} else {
// 				// x = nums[i]
// 				// f[i] 表示 选择在前 i 个数的范围内选择，
// 				// 即 在 数组下标 [0,i-1]的范围内选择
// 				// f[i_0] 中的 i_0 表示前 i 个数字
// 				// 而 nums[i_1] 中的 i_1 表示 数组下标，
// 				// 两者含义不一样 i_1 + 1 == i_0
// 				f[i+1][c] = f[i][c] + f[i][c-x] // 不选 + 选
// 			}
// 		}
// 	}
// 	// 在前n个数字的范围内选择，即在下标范围 [0,n-1] 这个n个数字中选择
// 	return f[n][m]
// }

// DP 优化 二维数组
// func findTargetSumWays(nums []int, target int) int {
// 	abs := func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}
// 		return x
// 	}
// 	s := 0
// 	for _, x := range nums {
// 		s += x
// 	}
// 	s -= abs(target)
// 	if s < 0 || s%2 == 1 {
// 		return 0
// 	}
// 	m := s / 2 // 背包容量

// 	n := len(nums)
// 	f := make([][]int, 2)
// 	for i := range f {
// 		f[i] = make([]int, m+1)
// 	}
// 	f[0][0] = 1
// 	for i, x := range nums {
// 		for c := 0; c <= m; c++ {
// 			if c < x {
// 				f[(i+1)%2][c] = f[i%2][c] // 只能不选
// 			} else {
// 				f[(i+1)%2][c] = f[i%2][c] + f[i%2][c-x] // 不选 + 选
// 			}
// 		}
// 	}
// 	// 在前n个数字的范围内选择，即在下标范围 [0,n-1] 这个n个数字中选择
// 	return f[n%2][m]
// }

// DP 优化 一维数组
func findTargetSumWays(nums []int, target int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	s := 0
	for _, x := range nums {
		s += x
	}
	s -= abs(target)
	if s < 0 || s%2 == 1 {
		return 0
	}
	m := s / 2 // 背包容量

	f := make([]int, m+1)

	f[0] = 1
	for _, x := range nums {
		for c := m; c >= x; c-- {
			f[c] = f[c] + f[c-x] // 不选 + 选
		}
	}
	// 在前n个数字的范围内选择，即在下标范围 [0,n-1] 这个n个数字中选择
	return f[m]
}

// @lc code=end
