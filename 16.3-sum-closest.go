/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.76%)
 * Likes:    1548
 * Dislikes: 111
 * Total Accepted:    411.5K
 * Total Submissions: 900.1K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 * 
 * Example:
 * 
 * 
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 * 
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 * 
 * -4 -1 1 2      1
 */

// @lc code=start
import "math"
func threeSumClosest(nums []int, target int) int {
	diff := math.MaxInt64

	sort.Sort(sort.IntSlice(nums))
	for i := 2; i < len(nums); i++ {
		// fmt.Println(i)
		if i < len(nums)-1 && nums[i+1] == nums[i] {
			continue
		}
		l := 0
		r := i-1

		for l < r {
			tmpDiff := nums[l] + nums[r] + nums[i] - target
			// fmt.Println(l, r, tmpDiff)
			if tmpDiff < 0 {
				lv := nums[l]
				for l < r && nums[l] == lv {
					l++
				}
			} else if tmpDiff > 0 {
				rv := nums[r]
				for l < r && nums[r] == rv {
					r--
				}
			} else {
				return target
			}
			if math.Abs(float64(tmpDiff)) < math.Abs(float64(diff)) {
				diff = tmpDiff
			}
		}
	}
	
	return target + diff
}
// @lc code=end

