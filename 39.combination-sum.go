/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (50.34%)
 * Likes:    2884
 * Dislikes: 90
 * Total Accepted:    450.4K
 * Total Submissions: 853.5K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 * 
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
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
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 * 
 * 
 */

// @lc code=start
func combinationSumImpl(candidates []int, target int) [][]int {
	var res [][]int
	if target <= 0 && len(candidates) <= 0 {
		return res
	}

	for i := 0; i < len(candidates); i++ {
		num := candidates[i]
		tmp := []int{num}
		sum := num
		for sum < target {
			tmpRes := combinationSumImpl(candidates[i+1:], target-sum)
			for _, v := range tmpRes {
				res = append(res, append(v, tmp...))	// v后续只有res引用，不会受到append影响
			}
			
			sum += num
			tmp = append(tmp, num)
		}
		if sum == target {
			res = append(res, tmp)
		}
	}
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	res := combinationSumImpl(candidates, target)
	return res
}
// @lc code=end

