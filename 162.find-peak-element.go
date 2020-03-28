/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 *
 * https://leetcode.com/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (42.63%)
 * Likes:    1365
 * Dislikes: 1818
 * Total Accepted:    318.8K
 * Total Submissions: 746.7K
 * Testcase Example:  '[1,2,3,1]'
 *
 * A peak element is an element that is greater than its neighbors.
 * 
 * Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element
 * and return its index.
 * 
 * The array may contain multiple peaks, in that case return the index to any
 * one of the peaks is fine.
 * 
 * You may imagine that nums[-1] = nums[n] = -∞.
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,2,3,1]
 * Output: 2
 * Explanation: 3 is a peak element and your function should return the index
 * number 2.
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [1,2,1,3,5,6,4]
 * Output: 1 or 5 
 * Explanation: Your function can return either index number 1 where the peak
 * element is 2, 
 * or index number 5 where the peak element is 6.
 * 
 * 
 * Note:
 * 
 * Your solution should be in logarithmic complexity.
 * 
 */
 package main
 import "fmt"

// @lc code=start
func findPeakElement(nums []int) int {
    if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		
		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			} else {
				lo = mid + 1
			}
		} else if mid == len(nums)-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			} else {
				hi = mid - 1
			}
		} else {
			if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
				return mid
			}
			if nums[mid] < nums[mid+1] {
				lo = mid+1
			} else {
				hi = mid-1
			}
		}
	}
	return -1
}
// @lc code=end

func main() {
	input := []int{2,1}
	res := findPeakElement(input)
	fmt.Println(res)
}
