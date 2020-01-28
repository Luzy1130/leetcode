/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 *
 * https://leetcode.com/problems/power-of-three/description/
 *
 * algorithms
 * Easy (41.81%)
 * Likes:    389
 * Dislikes: 1278
 * Total Accepted:    222K
 * Total Submissions: 529.4K
 * Testcase Example:  '27'
 *
 * Given an integer, write a function to determine if it is a power of three.
 * 
 * Example 1:
 * 
 * 
 * Input: 27
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 0
 * Output: false
 * 
 * Example 3:
 * 
 * 
 * Input: 9
 * Output: true
 * 
 * Example 4:
 * 
 * 
 * Input: 45
 * Output: false
 * 
 * Follow up:
 * Could you do it without using any loop / recursion?
 */

// @lc code=start
// package main
// import (
// 	"fmt"
// 	"math"
// )
func method1(n int) bool {
	// 3^x = n
	// xlog3 = logn
	// x = logn / log
	// 看x是否是整数即可，但是这里的复杂度并不是O(1)， 这里主要看计算Log10的复杂度
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	x := math.Log10(float64(n)) / math.Log10(float64(3))
	// fmt.Println(x)
	if (x - math.Floor(x)) < 0.0000000000001 {
		return true
	} else {
		return false
	}
}

func method2(n int) bool {
	// 1162261467是32位有符号数中3的最大幂次数
	return n > 0 && 1162261467 % n == 0;
}

func isPowerOfThree(n int) bool {
	return method2(n)
}

// func main() {
// 	fmt.Println(isPowerOfThree(2187))
// }
// @lc code=end

