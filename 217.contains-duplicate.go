/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (53.08%)
 * Likes:    562
 * Dislikes: 616
 * Total Accepted:    441.6K
 * Total Submissions: 812.1K
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an array of integers, find if the array contains any duplicates.
 * 
 * Your function should return true if any value appears at least twice in the
 * array, and it should return false if every element is distinct.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3,1]
 * Output: true
 * 
 * Example 2:
 * 
 * 
 * Input: [1,2,3,4]
 * Output: false
 * 
 * Example 3:
 * 
 * 
 * Input: [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 * 
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			return true
		}
		numMap[v] = true
	}
	return false
}
// @lc code=end

