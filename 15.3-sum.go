/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (24.64%)
 * Likes:    5243
 * Dislikes: 636
 * Total Accepted:    739.7K
 * Total Submissions: 2.9M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 * 
 * Note:
 * 
 * The solution set must not contain duplicate triplets.
 * 
 * Example:
 * 
 * 
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 * 
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}

	// 先做排序
	sort.Sort(sort.IntSlice(nums))
	// fmt.Println(nums)

	// 第三个数移动，前两个数使用sum2的方式寻找
	for i := 2; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i+1] == nums[i] {
			continue
		}
		l := 0
		r := i - 1
		for l < r {
			for l < r && nums[l]+nums[r] > -nums[i] {
				r--
			}
			for l < r && nums[l]+nums[r] < -nums[i] {
				l++
			}
			if l <  r && nums[l] + nums[r] == -nums[i] {
				res = append(res, []int{nums[l], nums[r], nums[i]})
				lv := nums[l]
				rv := nums[r]
				for l < r && nums[l] == lv {
					l++
				}
				for l < r && nums[r] == rv {
					r--
				}
			}
		}
	}
	return res
}
// @lc code=end

