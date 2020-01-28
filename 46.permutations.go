/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (57.02%)
 * Likes:    2934
 * Dislikes: 91
 * Total Accepted:    499.6K
 * Total Submissions: 836.8K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a collection of distinct integers, return all possible permutations.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,3]
 * Output:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 * 
 */

// @lc code=start
func permuteImpl(nums []int, usedMap map[int]bool) [][] int {
	var res [][]int

	if len(nums) == 0 {
		return res
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := usedMap[i]; ok {
			continue
		}
		tmp := []int{nums[i]}
		usedMap[i] = true
		otherRes := permuteImpl(nums, usedMap)
		for _, v := range otherRes {
			res = append(res, append(tmp, v...))
		}
		if len(otherRes) == 0 {
			res = append(res, tmp)
		}
		delete(usedMap, i)
	}
	return res
}

func permute(nums []int) [][]int {
	usedMap := make(map[int]bool)

	res := permuteImpl(nums, usedMap)
	return res
}
// @lc code=end

