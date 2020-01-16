/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (34.09%)
 * Likes:    2412
 * Dislikes: 111
 * Total Accepted:    401.5K
 * Total Submissions: 1.2M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * If the target is not found in the array, return [-1, -1].
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 * 
 */

// @lc code=start
func searchRangeBound(nums []int, target int, left bool) int {
	lo := 0
	hi := len(nums)

	for lo < hi {
		mid := (lo + hi) / 2
		// fmt.Println(lo, hi, mid)
		if nums[mid] == target {
			if left && (mid == 0 || nums[mid-1] != target) {
				return mid
			}
			if !left && (mid == len(nums)-1 || nums[mid+1] != target) {
				return mid
			}
		}
		if (left && nums[mid] == target) || nums[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	var res []int = []int{-1, -1}
	
	l := searchRangeBound(nums, target, true)
	if l == -1 {
		return res
	} else if l == len(nums)-1 {
		res[0] = l
		res[1] = l
	} else {
		r := searchRangeBound(nums, target, false)
		res[0] = l
		res[1] = r
	}
	return res
}
// @lc code=end

