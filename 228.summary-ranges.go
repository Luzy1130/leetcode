/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 *
 * https://leetcode.com/problems/summary-ranges/description/
 *
 * algorithms
 * Medium (38.30%)
 * Likes:    553
 * Dislikes: 486
 * Total Accepted:    157.9K
 * Total Submissions: 410.1K
 * Testcase Example:  '[0,1,2,4,5,7]'
 *
 * Given a sorted integer array without duplicates, return the summary of its
 * ranges.
 * 
 * Example 1:
 * 
 * 
 * Input:  [0,1,2,4,5,7]
 * Output: ["0->2","4->5","7"]
 * Explanation: 0,1,2 form a continuous range; 4,5 form a continuous range.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:  [0,2,3,4,6,8,9]
 * Output: ["0","2->4","6","8->9"]
 * Explanation: 2,3,4 form a continuous range; 8,9 form a continuous range.
 * 
 * 
 */
package main
import  (
	"fmt"
	"strconv"
)

// @lc code=start
func summaryRanges(nums []int) []string {
	var res []string
	if len(nums) == 0 {
		return res
	}
	
	i := 0
	j := 0
	for i < len(nums) {
		var tmp string = strconv.Itoa(nums[i])
		j++
		for j < len(nums) {
			if nums[j]-nums[j-1] != 1{
				break
			}
			j++
		}
		if j - i > 1 {
			tmp = tmp + "->" + strconv.Itoa(nums[j-1])
		}
		res = append(res, tmp)
		i = j
	}
	return res
}
// @lc code=end

func main() {
	res := summaryRanges([]int{0,2,3,4,6,8,9})
	fmt.Println(res)
}