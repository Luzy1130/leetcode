/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 *
 * https://leetcode.com/problems/majority-element-ii/description/
 *
 * algorithms
 * Medium (32.90%)
 * Likes:    1264
 * Dislikes: 147
 * Total Accepted:    126.6K
 * Total Submissions: 369.7K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an integer array of size n, find all elements that appear more than ⌊
 * n/3 ⌋ times.
 * 
 * Note: The algorithm should run in linear time and in O(1) space.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,2,3]
 * Output: [3]
 * 
 * Example 2:
 * 
 * 
 * Input: [1,1,1,3,3,2,2,2]
 * Output: [1,2]
 * 
 */
// 还有一种随机的算法也蛮有意思的，但是不能保证100%准确


// https://leetcode.com/problems/majority-element-ii/discuss/543672/BoyerMoore-majority-vote-algorithm-EXPLAINED-(with-pictures)
// 数组中最多有2个符合条件的元素(1/3 * 2)
// 投票算法（vote）， 设置两个投票箱

package main
import "fmt"

// @lc code=start
func majorityElement(nums []int) []int {
	s1 := -1
	s2 := -1
	v1 := 0
	v2 := 0
	for _, num := range nums {
		if s1 == num {
			v1++
		} else if s2 == num {
			v2++
		} else if v1 == 0 {
			s1 = num
			v1++
		} else if v2 == 0 {
			s2 = num
			v2++
		} else {
			v1--
			v2--
		}
	}

	v1 = 0
	v2 = 0
	for _, num := range nums {
		if num == s1 {
			v1++
		} else if num == s2 {
			v2++
		}
	}

	var res []int
	if v1 > len(nums)/3 {
		res = append(res, s1)
	}
	if v2 > len(nums)/3 {
		res = append(res, s2)
	}
	return res
}
// @lc code=end

func main() {
	res := majorityElement([]int{1,1,1,3,3,2,2,2})
	fmt.Println(res)

}
