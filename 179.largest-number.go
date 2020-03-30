/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 *
 * https://leetcode.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (27.55%)
 * Likes:    1561
 * Dislikes: 188
 * Total Accepted:    163.9K
 * Total Submissions: 592.8K
 * Testcase Example:  '[10,2]'
 *
 * Given a list of non negative integers, arrange them such that they form the
 * largest number.
 * 
 * Example 1:
 * 
 * 
 * Input: [10,2]
 * Output: "210"
 * 
 * Example 2:
 * 
 * 
 * Input: [3,30,34,5,9]
 * Output: "9534330"
 * 
 * 
 * Note: The result may be very large, so you need to return a string instead
 * of an integer.
 * 
 */
package main
import (
	"fmt"
	"sort"
	"strconv"
)

// @lc code=start

type intSlice []int

func (s intSlice) Len() int { return len(s) }
func (s intSlice) Less(i, j int) bool {
	stri := []rune(strconv.Itoa(s[i]))
	strj := []rune(strconv.Itoa(s[j]))

	// 这里怎么排序是关键
	for len(stri) != 0 || len(strj) != 0 {
		k := 0
		for ; k < len(stri) && k < len(strj); k++ {
			if rune(stri[k]) > rune(strj[k]) {
				return true
			}
			if rune(stri[k]) < rune(strj[k]) {
				return false
			}
		}
		if k < len(stri) {
			stri = stri[k:]
		} else if k < len(strj) {
			strj = strj[k:]
		} else {
			break
		}
	}

	return false
	// // return stri[len(stri)-1] > strj[len(strj)-1]
	// if k < len(stri) {
	// 	return stri[len(stri)-1] > stri[0]
	// } else {
	// 	return strj[len(strj)-1] <= strj[0]
	// }
}
func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func largestNumber(nums []int) string {
	sort.Sort(intSlice(nums))
	// fmt.Println(nums)

	nonzeroCnt := 0
	var res string
	for i := 0; i < len(nums); i++ {
		res = res + strconv.Itoa(nums[i])
		if nums[i] > 0 {
			nonzeroCnt++
		}
	}

	if nonzeroCnt == 0 {
		return "0"
	}
	return res
}
// @lc code=end

func main() {
	res := largestNumber([]int{1,1,1})
	fmt.Println(res)
}