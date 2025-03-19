package go_case

import "sort"

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
type SortedMap struct {
	toes   []string       // 有序键序列
	values map[string]int // 实际键值对
}

func NewSortedMap() *SortedMap {
	return &SortedMap{
		toes:   []string{},
		values: make(map[string]int),
	}
}

func (s *SortedMap) Set(to string) {
	if _, ok := s.values[to]; !ok {
		s.toes = append(s.toes, to)
	}
	s.values[to]++
}

func (s *SortedMap) Sort() {
	sort.Sort(sort.StringSlice(s.toes))
	return
}

func findItinerary(tickets [][]string) []string {

	targetsMap := make(map[string]*SortedMap)
	for _, pair := range tickets {
		from, to := pair[0], pair[1]
		if _, ok := targetsMap[from]; !ok {
			targetsMap[from] = NewSortedMap()
		}
		targetsMap[from].Set(to)
	}
	for k, _ := range targetsMap {
		targetsMap[k].Sort()
	}

	res := []string{"JFK"}
	var bt func() bool
	bt = func() bool {
		if len(res) == len(tickets)+1 {
			return true
		}
		start := res[len(res)-1]
		destinations := targetsMap[start]
		if destinations != nil {
			for _, to := range destinations.toes {
				if destinations.values[to] > 0 {
					destinations.values[to]--
					res = append(res, to)
					if bt() {
						return true
					}
					res = res[:len(res)-1]
					destinations.values[to]++
				}
			}
		}
		return false
	}
	bt()
	return res
}

// @lc code=end
