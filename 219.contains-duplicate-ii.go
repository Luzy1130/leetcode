/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 *
 * https://leetcode.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (35.86%)
 * Likes:    646
 * Dislikes: 792
 * Total Accepted:    232.8K
 * Total Submissions: 637.5K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * Given an array of integers and an integer k, find out whether there are two
 * distinct indices i and j in the array such that nums[i] = nums[j] and the
 * absolute difference between i and j is at most k.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,2,3,1], k = 3
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [1,0,1,1], k = 1
 * Output: true
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: nums = [1,2,3,1,2,3], k = 2
 * Output: false
 * 
 * 
 * 
 * 
 * 
 */
// @lc code=start

func method1(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] == nums[j] && j-i <= k {
				return true
			}
		}
	}
	return false
}

func method2(nums []int, k int) bool {
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v, ok := numMap[nums[i]]; ok && i-v<=k {
			return true
		} else {
			numMap[nums[i]] = i
		}
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	return method2(nums, k)
}
// @lc code=end

