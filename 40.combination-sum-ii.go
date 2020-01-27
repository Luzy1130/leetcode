/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (43.14%)
 * Likes:    1284
 * Dislikes: 54
 * Total Accepted:    279.2K
 * Total Submissions: 618.1K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sums to target.
 * 
 * Each number in candidates may only be used once in the combination.
 * 
 * Note:
 * 
 * 
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 * [1,2,2],
 * [5]
 * ]
 * 
 * 
 */

// @lc code=start
func combinationSumImpl(candidates []int, target int) [][]int {
	var res [][]int
	if target <= 0 {
		return res
	}

	for i := 0; i < len(candidates); i++ {
		num := candidates[i]
		tmp := []int{num}

		if num < target {
			resTmp := combinationSumImpl(candidates[i+1:], target-num)
			for _, v := range resTmp {
				res = append(res, append(tmp, v...))
			} 
		} else if num == target {
			res = append(res, tmp)
		}

		// 滑过重复的数字
		for i < len(candidates)-1 && candidates[i+1] == candidates[i] {
			i++
		}
	}
	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	res := combinationSumImpl(candidates, target)
	return res
}
// @lc code=end

