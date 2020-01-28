/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (41.88%)
 * Likes:    1509
 * Dislikes: 52
 * Total Accepted:    304.2K
 * Total Submissions: 694K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers that might contain duplicates, return all
 * possible unique permutations.
 * 
 * Example:
 * 
 * 
 * Input: [1,1,2]
 * Output:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 * 
 * 
 */

// @lc code=start
func permuteUniqueImpl(nums []int, usedMap map[int]bool) [][]int {
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
		otherRes := permuteUniqueImpl(nums, usedMap)
		for _, v := range otherRes {
			res = append(res, append(tmp, v...))
		}
		if len(otherRes) == 0 {
			res = append(res, tmp)
		}
		delete(usedMap, i)

		// 滑动，略过相同的元素
		for i < len(nums)-1 && nums[i+1] == nums[i] {
			i++
		}
	}
	return res

}

func permuteUnique(nums []int) [][]int {
	usedMap := make(map[int]bool)

	sort.Sort(sort.IntSlice(nums))
	res := permuteUniqueImpl(nums, usedMap)
	
	return res
}
// @lc code=end

