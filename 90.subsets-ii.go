/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (43.46%)
 * Likes:    1364
 * Dislikes: 59
 * Total Accepted:    252.4K
 * Total Submissions: 556.2K
 * Testcase Example:  '[1,2,2]'
 *
 * Given a collection of integers that might contain duplicates, nums, return
 * all possible subsets (the power set).
 * 
 * Note: The solution set must not contain duplicate subsets.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,2]
 * Output:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 * 
 * 
 */

// @lc code=start
// package main
import (
	"encoding/json"
	"sort"
	// "fmt"
)


func subSetWithDupImpl(nums []int, hash map[string]bool) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}

	firstNum := nums[0]
	suffixSets := subsetsWithDup(nums[1:])

	res := [][]int{}

	for _, child := range suffixSets {
		newChild := []int{firstNum}
		newChild = append(newChild, child...)
		childStr, _ := json.Marshal(newChild)
		if _, ok := hash[string(childStr)]; ok {
			continue
		}
		hash[string(childStr)] = true
		res = append(res, newChild)
	}
	for _, child := range suffixSets {
		childStr, _ := json.Marshal(child)
		if _, ok := hash[string(childStr)]; ok {
			continue
		}
		hash[string(childStr)] = true
		res = append(res, child)
	}
	return res
}

// 关键在于怎么去重
func subsetsWithDup(nums []int) [][]int {
	hash := make(map[string]bool)
	// 排序是因为，两个数组数据相同顺序不同，也算是同一个元素
	sort.Sort(sort.IntSlice(nums))
	res := subSetWithDupImpl(nums, hash)
	return res
}

// func main() {
// 	res := subsetsWithDup([]int{1,2,2,1})
// 	fmt.Println(res)

// }
// @lc code=end

