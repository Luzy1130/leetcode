/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 *
 * https://leetcode.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (39.41%)
 * Likes:    662
 * Dislikes: 912
 * Total Accepted:    170.4K
 * Total Submissions: 412K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 * 
 * Example:
 * 
 * Given nums = [-2, 0, 3, -5, 2, -1]
 * 
 * sumRange(0, 2) -> 1
 * sumRange(2, 5) -> -1
 * sumRange(0, 5) -> -3
 * 
 * 
 * 
 * Note:
 * 
 * You may assume that the array does not change.
 * There are many calls to sumRange function.
 * 
 * 
 */

// @lc code=start
type NumArray struct {
    Sums []int
}


func Constructor(nums []int) NumArray {
	arr := NumArray{}
	// arr.Sums = append(arr.Sums, 0)
	sum := 0
	for _, v := range nums {
		sum += v
		arr.Sums = append(arr.Sums, sum)
	}
	return arr
}


func (this *NumArray) SumRange(i int, j int) int {
	sums := this.Sums
	
	var res = sums[j]
	if i != 0 {
		res -= sums[i-1]
	}
	return res
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

