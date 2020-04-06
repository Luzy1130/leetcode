/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 *
 * https://leetcode.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (20.51%)
 * Likes:    929
 * Dislikes: 970
 * Total Accepted:    114.8K
 * Total Submissions: 557.8K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * Given an array of integers, find out whether there are two distinct indices
 * i and j in the array such that the absolute difference between nums[i] and
 * nums[j] is at most t and the absolute difference between i and j is at most
 * k.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,2,3,1], k = 3, t = 0
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [1,0,1,1], k = 1, t = 2
 * Output: true
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: nums = [1,5,9,1,5,9], k = 2, t = 3
 * Output: false
 * 
 * 
 * 
 * 
 */
package main
import (
	"fmt"
	"sort"
)

// @lc code=start
type pair struct {
	val int
	index int
}
type pairs []pair

func (p pairs) Len() int {return len(p)}
func (p pairs) Less(i, j int) bool { return p[i].val < p[j].val}
func (p pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	ps := make(pairs, len(nums))
	for i := 0; i < len(nums); i++ {
		ps[i].val = nums[i]
		ps[i].index = i
	}
	sort.Sort(ps)
	// fmt.Println(ps)

	for i := 0; i < len(ps)-1; i++ {
		for j := i+1; j < len(ps);j++ {
			numD := ps[j].val - ps[i].val
			indexD := ps[j].index - ps[i].index
			if numD <0 {
				numD = -numD
			}
			if indexD < 0 {
				indexD = -indexD
			}

			if numD <= t && indexD <= k {
				return true
			}
			if numD > t {
				break
			}
		}
	}
	return false
}
// @lc code=end

func main() {
	res := containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9}, 2, 3)
	// res := containsNearbyAlmostDuplicate([]int{1,3,6,2}, 1, 2)
	fmt.Println(res)
}