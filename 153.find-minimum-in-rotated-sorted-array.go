/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 *
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (44.24%)
 * Likes:    1675
 * Dislikes: 214
 * Total Accepted:    387.8K
 * Total Submissions: 875.9K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
 * 
 * Find the minimum element.
 * 
 * You may assume no duplicate exists in the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,4,5,1,2] 
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [4,5,6,7,0,1,2]
 * Output: 0
 * 
 * 
 */
package main
import "fmt"

// @lc code=start

func findMin(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	low := 0
	high := len(nums)

	// 如果是有序的直接返回即可
	if nums[low] <= nums[high-1] {
		return nums[low]
	}

	for low < high {
		mid := (low+high) / 2
		// fmt.Println(low, high, mid)

		if mid > low && mid < high-1 {
			// 如果mid两边有数据
			if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
				return nums[mid+1]
			}
			if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
				return nums[mid]
			}
			if nums[mid] < nums[low] {
				high = mid
			} else {
				low = mid + 1
			}
		} else {
			// mid如果在边界上，通过简单判断后也可直接返回
			if mid == low {
				if nums[mid] <= nums[high-1] {
					return nums[mid]
				} else {
					return nums[high-1]
				}
			} else {
				if nums[mid] <= nums[low] {
					return nums[mid]
				} else {
					return nums[low]
				}
			}
		}
	}

	return nums[low]
}
// @lc code=end


func main() {
	input := []int{ 1,2,3}
	res := findMin(input)
	fmt.Println(res)
}
