/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (49.40%)
 * Likes:    1135
 * Dislikes: 59
 * Total Accepted:    253.2K
 * Total Submissions: 489.1K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 * 
 * Example:
 * 
 * 
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"
func combineImpl(h, l, k int) [][]int {
	var res [][]int
	if k == 0 {
		return res
	}

	for i := l; i <= h; i++ {
		var tmp []int = []int{i}
		if k == 1{
			res = append(res, tmp)
		} else {
			suffix := combineImpl(h, i+1, k-1)
			for _, v := range suffix {
				res = append(res, append(tmp, v...))
			}
		}
	}
	return res
}

func combine(n int, k int) [][]int {
	res := combineImpl(n, 1, k)
	return res
}

// func main() {
// 	res := combine(4,3)
// 	fmt.Println(res)
// }
// @lc code=end

