package go_case

/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 *
 * https://leetcode.cn/problems/reconstruct-itinerary/description/
 *
 * algorithms
 * Hard (45.84%)
 * Likes:    896
 * Dislikes: 0
 * Total Accepted:    103.1K
 * Total Submissions: 225.1K
 * Testcase Example:  '[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]'
 *
 * 给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi]
 * 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。
 *
 * 所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK
 * 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。
 *
 *
 * 例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
 *
 *
 * 假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
 * 输出：["JFK","MUC","LHR","SFO","SJC"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：tickets =
 * [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
 * 输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
 * 解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * tickets[i].length == 2
 * fromi.length == 3
 * toi.length == 3
 * fromi 和 toi 由大写英文字母组成
 * fromi != toi
 *
 *
 */

// @lc code=start
func findItinerary(tickets [][]string) []string {

	var res []string
	var path []string
	ticketUsed := make([]bool, len(tickets))
	for i, _ := range ticketUsed {
		ticketUsed[i] = false
	}

	var backtracing func(string) bool

	backtracing = func(cur string) bool {
		if len(path) == len(tickets) {
			res = append([]string{}, path...)
			res = append([]string{}, cur)
			return true
		}
		for i, ticket := range tickets {
			// if ticket 被使用过 则continue
			if ticketUsed[i] {
				continue
			}
			if ticket[0] == cur {
				path = append(path, ticket[0])
				ticketUsed[i] = true
				if backtracing(ticket[1]) {
					return true
				}
				ticketUsed[i] = false
				path = path[:len(path)-1]

			}

		}
		return false

	}

	backtracing("JFK")
	return res

}

// @lc code=end
