/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 *
 * https://leetcode.com/problems/single-number-iii/description/
 *
 * algorithms
 * Medium (59.64%)
 * Likes:    1262
 * Dislikes: 94
 * Total Accepted:    130.6K
 * Total Submissions: 217.3K
 * Testcase Example:  '[1,2,1,3,2,5]'
 *
 * Given an array of numbers nums, in which exactly two elements appear only
 * once and all the other elements appear exactly twice. Find the two elements
 * that appear only once.
 * 
 * Example:
 * 
 * 
 * Input:  [1,2,1,3,2,5]
 * Output: [3,5]
 * 
 * Note:
 * 
 * 
 * The order of the result is not important. So in the above example, [5, 3] is
 * also correct.
 * Your algorithm should run in linear runtime complexity. Could you implement
 * it using only constant space complexity?
 * 
 */
package main
import "fmt"

// @lc code=start
func singleNumber(nums []int) []int {
	mask := 0
	for _, v := range nums {
		mask ^= v
	}

	// 两个数字至少有一个bit是不同的,mask中不为0的bit即是两个数字不同的bit
	// A和-A与操作得到A中最小的bit为1的数字
	// filter代表两个数有区别的最小bit
	filter := mask & (-mask)
	fmt.Printf("%b,%b,%b\n", mask, -mask, filter)
	var res []int = []int{0, 0}
	for _, v := range nums {
		if (filter&v) != 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
// @lc code=end

func main() {
	res := singleNumber([]int{2,1,2,3,4,1})
	fmt.Println(res)
}