/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (33.11%)
 * Likes:    3517
 * Dislikes: 391
 * Total Accepted:    549.2K
 * Total Submissions: 1.6M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 * 
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 * 
 * You may assume no duplicate exists in the array.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * 
 */

// @lc code=start
// package main

// import "fmt"

//                         
// 原来有序数组只有递增的关系  /
// 经过旋转后，数组出现了两个不连续的递增关系 / / 
// 且第二个递增的数组全部小于第一个递增的数组，可画一下图，来想象下一个查找区间是左还是右
func search(nums []int, target int) int {
	lo := 0
	hi := len(nums)

	for lo < hi {
		mid := (lo + hi) / 2
		// 如果找到目标则直接返回
		if nums[mid] == target {
			return mid
		}

		// 如果没找到目标，则需要确定继续在左还是右区间找
		if nums[mid] > target {
			// 中间数比目标数大，则正常的有序数组中就应该在左区间找，
			// 1. 如果log比mid大，且mid比target大，则target在lo和mid之间；
			// 2. 如果target比lo大，且比mid小，则target在lo和mid之间
			if nums[lo] > nums[mid] || nums[lo] <= target {
				hi = mid
			} else {
				lo = mid + 1
			}
		} else {
			if nums[hi-1] < nums[mid] || nums[hi-1] >= target {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	}

	return -1
}

// func main() {
// 	res := search([]int{4,5,6,7,0,1,2}, 0)
// 	fmt.Println(res)
// }
// @lc code=end

