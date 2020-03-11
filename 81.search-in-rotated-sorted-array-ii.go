/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (32.75%)
 * Likes:    955
 * Dislikes: 401
 * Total Accepted:    208.9K
 * Total Submissions: 634.7K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
 * 
 * You are given a target value to search. If found in the array return true,
 * otherwise return false.
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [2,5,6,0,0,1,2], target = 0
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [2,5,6,0,0,1,2], target = 3
 * Output: false
 * 
 * Follow up:
 * 
 * 
 * This is a follow up problem to Search in Rotated Sorted Array, where nums
 * may contain duplicates.
 * Would this affect the run-time complexity? How and why?
 * 
 * 
 */

// @lc code=start
// package main

// import "fmt"

func search(nums []int, target int) bool {
	lo := 0
	hi := len(nums)

	for lo < hi {
		mid := (lo + hi) / 2

		if target == nums[mid] {
			return true
		}

		if target < nums[mid] {
			if nums[lo] == nums[mid] {
				lo++
			} else if (nums[lo] > nums[mid] || target >= nums[lo]) {
				hi = mid
			} else {
				lo = mid + 1
			}
		} else {
			if nums[hi-1] == nums[mid] {
				hi--
			} else if (nums[hi-1] < nums[mid] || target <= nums[hi-1]) {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	}

	return false
}


// func main() {
// 	res := search([]int{3,1,1}, 3)
// 	fmt.Println(res)
// }
// @lc code=end

