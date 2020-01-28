/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 *
 * https://leetcode.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (40.64%)
 * Likes:    397
 * Dislikes: 173
 * Total Accepted:    133.4K
 * Total Submissions: 324.6K
 * Testcase Example:  '16'
 *
 * Given an integer (signed 32 bits), write a function to check whether it is a
 * power of 4.
 * 
 * Example 1:
 * 
 * 
 * Input: 16
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 5
 * Output: false
 * 
 * 
 * Follow up: Could you solve it without loops/recursion?
 */

// @lc code=start
// package main
// import (
// 	"fmt"
// )
func isPowerOfFour(num int) bool {
    if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}
	
	// 通过移位的话需要循环，可题目限制不符
	cnt := 0
	for num != 0 {
		if num & 0x1 == 0 {
			cnt++
		} else {
			break
		}
		num = num >> 1
	}
	if num == 1 && cnt % 2 == 0 {
		return true
	} else {
		return false
	}
}

// func main() {
// 	fmt.Println(isPowerOfFour(64))
// }
// @lc code=end

