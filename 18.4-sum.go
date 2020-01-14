/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (31.35%)
 * Likes:    1448
 * Dislikes: 274
 * Total Accepted:    286.7K
 * Total Submissions: 889.9K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 * 
 * Note:
 * 
 * The solution set must not contain duplicate quadruplets.
 * 
 * Example:
 * 
 * 
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 * 
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 * 
 * 
 */

// @lc code=start

func kSum(nums[]int, target int, k int) [][]int {
	var res [][]int
	if k <= 0 || len(nums) < k {
		return res
	}

	if k == 1 {
		// fmt.Println(nums, target, k)
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				res = append(res, []int{nums[i]})
				return res
			}
		}
		return res
	}

	for i := 0; i < len(nums)-k+1; {
		newnums := nums[i+1:]
		// fmt.Println(i, nums)
		tmp := kSum(newnums, target-nums[i], k-1)
		for j := 0; j < len(tmp); j++ {
			res = append(res, append([]int{nums[i]}, tmp[j]...))
		}
		
		for i < len(nums)-k+1 {
			i++
			if nums[i-1] != nums[i] {
				break
			}
		}
	}
	return res
}

// 深度遍历算法
func dfs(nums []int, target int) [][]int {
	sort.Sort(sort.IntSlice(nums))

	res := kSum(nums, target, 4)
	return res
}

// 两个指针的算法
func twoPointer(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < len(nums)-3; {
		for j := i+1; j < len(nums)-2; {
			newTarget := target-nums[i]-nums[j]
			l := j+1
			r := len(nums)-1
			for l < r {
				lv := nums[l]
				rv := nums[r]
				if nums[l]+nums[r] < newTarget {
					for l < r && nums[l] == lv {
						l++
					}
				} else if nums[l]+nums[r] > newTarget {
					for l < r && nums[r] == rv {
						r--
					}
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == lv {
						l++
					}
					for l < r && nums[r] == rv {
						r--
					}
				}
			}
			for j < len(nums)-2 {
				j++
				if nums[j-1] != nums[j] {
					break
				}
			}
		}
		for i < len(nums)-3 {
			i++
			if nums[i-1] != nums[i] {
				break
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	// return dfs(nums, target)
	return twoPointer(nums, target)
}
// @lc code=end

