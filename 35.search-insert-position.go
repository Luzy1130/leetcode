/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (41.14%)
 * Likes:    1504
 * Dislikes: 197
 * Total Accepted:    446.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array and a target value, return the index if the target is
 * found. If not, return the index where it would be if it were inserted in
 * order.
 * 
 * You may assume no duplicates in the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,3,5,6], 5
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [1,3,5,6], 2
 * Output: 1
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: [1,3,5,6], 7
 * Output: 4
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: [1,3,5,6], 0
 * Output: 0
 * 
 * 
 */
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums)
	i := (low + high) / 2
	
	if len(nums) == 0 {
		return 0;
	}

	// use binary-search
	for ;; {
		if i == low && nums[i] >= target{
			return low
		}
		if i == high - 1 && nums[i] < target {
			return high
		}
	
		if nums[i-1] < target && nums[i] >= target {
			return i
		}

		if nums[i] < target {
			low = i
		} else {
			high = i
		}
		i = (low + high) / 2
	}
}

