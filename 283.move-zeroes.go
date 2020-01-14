/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 *
 * https://leetcode.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (54.99%)
 * Likes:    2747
 * Dislikes: 97
 * Total Accepted:    587.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * Given an array nums, write a function to move all 0's to the end of it while
 * maintaining the relative order of the non-zero elements.
 * 
 * Example:
 * 
 * 
 * Input: [0,1,0,3,12] [1,0,0,3,12] [1,3,0,0,12] [1,3,12,0,0]
 * Output: [1,3,12,0,0]
 * 
 * Note:
 * 
 * 
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 * 
 */

// @lc code=start
// 一般这种挪位置的问题，用元素交换的方式去想想
func moveZeroes(nums []int)  {
	zeroL := 0
	zeroR := 0

	findZero := false
	// 先找到初始的连续0的范围
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && findZero == false {
			zeroL = i
			findZero = true
		}
		if nums[i] != 0 && findZero == true {
			zeroR = i-1
			break
		}
	}
	if findZero == false {
		return
	}

	for i := zeroR + 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroL], nums[i] = nums[i], nums[zeroL]
			zeroL++
		} 
		zeroR = i
	}
	return
}
// @lc code=end

