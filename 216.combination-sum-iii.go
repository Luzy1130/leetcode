/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 *
 * https://leetcode.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (54.72%)
 * Likes:    889
 * Dislikes: 47
 * Total Accepted:    153.5K
 * Total Submissions: 279.4K
 * Testcase Example:  '3\n7'
 *
 * 
 * Find all possible combinations of k numbers that add up to a number n, given
 * that only numbers from 1 to 9 can be used and each combination should be a
 * unique set of numbers.
 * 
 * Note:
 * 
 * 
 * All numbers will be positive integers.
 * The solution set must not contain duplicate combinations.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: k = 3, n = 7
 * Output: [[1,2,4]]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: k = 3, n = 9
 * Output: [[1,2,6], [1,3,5], [2,3,4]]
 * 
 * 
 */
package main
import "fmt"

// @lc code=start
func combinationSumImpl(k, n int, min int) [][]int {
	if k == 0 || n == 0 {
		return [][]int{}
	}

	var res [][]int
	for i := min; i <= 9 && i <= n; i++ {
		tmp := []int{i}
		suffixes := combinationSumImpl(k-1, n-i, i+1)
		if len(suffixes) != 0 {
			for _, suffix := range suffixes {
				if len(suffix) == k - 1 {
					tmpRes := append(tmp, suffix...)
					res = append(res, tmpRes)
				}
			}
		} else {
			if k == 1 && i == n {
				res = append(res, tmp)
			}
		}
	}

	return res
}

func combinationSum3(k int, n int) [][]int {
	res := combinationSumImpl(k, n, 1)
	return res
}
// @lc code=end

func main() {
	res :=  combinationSum3(3, 9)
	fmt.Println(res)
}