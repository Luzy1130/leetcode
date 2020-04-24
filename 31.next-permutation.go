/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (30.96%)
 * Likes:    2575
 * Dislikes: 867
 * Total Accepted:    306.4K
 * Total Submissions: 969.5K
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 * 
 * If such arrangement is not possible, it must rearrange it as the lowest
 * possible order (ie, sorted in ascending order).
 * 
 * The replacement must be in-place and use only constant extra memory.
 * 
 * Here are some examples. Inputs are in the left-hand column and its
 * corresponding outputs are in the right-hand column.
 * 
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 * 
 */
package main
import (
	"sort"
	"fmt"
)

// @lc code=start
func reverse(nums []int, start int) {
    i := start
    j := len(nums)-1
    for i < j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}

func nextPermutation(nums []int)  {
    i := len(nums) - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    if i >= 0 {
        j := len(nums)-1
        for j >= 0 && nums[j] <= nums[i] {
            j--
        }
        sort.IntSlice(nums).Swap(i, j)
    } 
    reverse(nums, i+1) 
}
// @lc code=end

func main() {
	input := []int{1,2,3}
	nextPermutation(input)
	fmt.Println(input)
}