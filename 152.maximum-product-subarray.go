/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (30.90%)
 * Likes:    3336
 * Dislikes: 139
 * Total Accepted:    295.5K
 * Total Submissions: 954.1K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find the contiguous subarray within an array
 * (containing at least one number) which has the largest product.
 * 
 * Example 1:
 * 
 * 
 * Input: [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 * 
 */

// @lc code=start
// package main
// import (
// 	"fmt"
// 	// "math"
// )

func maxProduct(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	curMax := nums[0]
	curMin := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		// fmt.Println(curMax, curMin)
		tmpMin := curMin * nums[i]
		tmpMax := curMax * nums[i]

		if tmpMin > tmpMax {
			tmpMax, tmpMin = tmpMin, tmpMax
		} 
		curMax = tmpMax
		curMin = tmpMin 

		// curMax = int(math.Max(math.Max(float64(tmpMin), float64(tmpMax)), float64(curMax)))
		// curMin = int(math.Min(math.Min(float64(tmpMin), float64(tmpMax)), float64(curMin)))

		if curMax < nums[i] {
			curMax = nums[i]
		}
		if curMin > nums[i] {
			curMin = nums[i]
		}

		if res < curMax {
			res  = curMax
		}
	}
	return res
}

// func main() {
// 	res := maxProduct([]int{-2, 0, -1})
// 	fmt.Println(res)
// }
// @lc code=end

