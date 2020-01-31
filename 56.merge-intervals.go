/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (36.44%)
 * Likes:    3138
 * Dislikes: 241
 * Total Accepted:    485.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given a collection of intervals, merge all overlapping intervals.
 * 
 * Example 1:
 * 
 * 
 * Input: [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 * 
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 * 
 */

// @lc code=start
// package main
// import (
// 	"fmt"
// 	"sort"
// )

type TwoIntSlice [][]int

func (s TwoIntSlice) Len() int { return len(s) }
func (s TwoIntSlice) Less(i, j int) bool { return s[i][0] < s[j][0] }
func (s TwoIntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 {
		return res
	}

	intervalsTmp := TwoIntSlice(intervals)
	sort.Sort(intervalsTmp)

	res = append(res, intervals[0])
	j := 0
	for i := 1; i < len(intervals); i++ {
		if res[j][1] >= intervals[i][0] {
			res[j][1] = int(math.Max(float64(res[j][1]),float64(intervals[i][1])))
		} else {
			res = append(res, intervals[i]) 
			j++
		}
	}
	return res
}

// func main() {
// 	var input [][]int = [][]int {
// 		[]int{1, 3},
// 		[]int{2, 6},
// 		[]int{8,10},
// 		[]int{15,18},
// 	}
// 	res := merge(input)
// 	fmt.Println(res)
// }
// @lc code=end

