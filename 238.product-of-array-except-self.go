/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (58.76%)
 * Likes:    3930
 * Dislikes: 335
 * Total Accepted:    416.1K
 * Total Submissions: 704K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an array nums of n integers where n > 1,  return an array output such
 * that output[i] is equal to the product of all the elements of nums except
 * nums[i].
 * 
 * Example:
 * 
 * 
 * Input:  [1,2,3,4]
 * Output: [24,12,8,6]
 * 
 * 
 * Constraint: It's guaranteed that the product of the elements of any prefix
 * or suffix of the array (including the whole array) fits in a 32 bit
 * integer.
 * 
 * Note: Please solve it without division and in O(n).
 * 
 * Follow up:
 * Could you solve it with constant space complexity? (The output array does
 * not count as extra space for the purpose of space complexity analysis.)
 * 
 */
package main
import "fmt"

// @lc code=start
func productExceptSelf(nums []int) []int {
   	var res []int = make([]int, len(nums))
	
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	tmp := 1
	for i := len(nums)-2; i >=0; i-- {
		tmp *= nums[i+1]
		res[i] = res[i] * tmp
	}

	return res
}
// @lc code=end

func main() {
	res := productExceptSelf([]int{1,2,3,4})
	fmt.Println(res)
}