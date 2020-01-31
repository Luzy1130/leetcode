/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 *
 * https://leetcode.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (32.49%)
 * Likes:    2929
 * Dislikes: 264
 * Total Accepted:    353K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers, you are initially positioned at the
 * first index of the array.
 * 
 * Each element in the array represents your maximum jump length at that
 * position.
 * 
 * Determine if you are able to reach the last index.
 * 
 * Example 1:
 * 
 * 
 * Input: [2,3,1,1,4]
 * Output: true
 * Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last
 * index.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,2,1,0,4]
 * Output: false
 * Explanation: You will always arrive at index 3 no matter what. Its
 * maximum
 * jump length is 0, which makes it impossible to reach the last index.
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

func canJump(nums []int) bool {
    if len(nums) == 0 {
		return true
	}

	// 从终点向前遍历，如果某个位置能跳到终点，则该位置可以认为和终点等价
	// 然后再判断前面的位置哪个能调到该等价位置
	// 如果最后0也是等价的位置，则说明可以跳到终点
	finalIdx := len(nums)-1
	for i := len(nums)-1; i >= 0; i-- {
		if i + nums[i] >= finalIdx {
			finalIdx = i
		}
	}

	return finalIdx == 0
}

// func main() {
// 	// var input []int = []int{2,3,1,1,4}
// 	// var input []int = []int{3,2,1,0,4}
// 	var input []int = []int{2,5,0,0}
// 	res := canJump(input)
// 	fmt.Println(res)
// }
// @lc code=end

