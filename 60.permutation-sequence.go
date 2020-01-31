/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 *
 * https://leetcode.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (33.94%)
 * Likes:    1143
 * Dislikes: 283
 * Total Accepted:    162.3K
 * Total Submissions: 462.7K
 * Testcase Example:  '3\n3'
 *
 * The set [1,2,3,...,n] contains a total of n! unique permutations.
 * 
 * By listing and labeling all of the permutations in order, we get the
 * following sequence for n = 3:
 * 
 * 
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 * 
 * 
 * Given n and k, return the k^th permutation sequence.
 * 
 * Note:
 * 
 * 
 * Given n will be between 1 and 9 inclusive.
 * Given k will be between 1 and n! inclusive.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: n = 3, k = 3
 * Output: "213"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: n = 4, k = 9
 * Output: "2314"
 * 
 * 
 */

// @lc code=start
// package main
// import (
// 	"fmt"
// 	"strconv"
// )

var facs []int = []int {
	1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880,
}

// l代表返回结果的长度
func getPermutationImpl(n, k, l int, usedMap map[int]bool) string {
	var res string
	if l == 0 {
		return res
	}

	a := (k-1) / facs[l-1] + 1	// a代表第几组
	b := (k-1) % facs[l-1] + 1	// b子序列第几个
	
	cnt := 0
	var prefix string
	for i := 1; i <= n; i++ {
		if _, ok := usedMap[i]; !ok {
			cnt++
			if cnt == a {
				prefix = strconv.Itoa(i)
				usedMap[i] = true
			}
		}
	}

	res = prefix + getPermutationImpl(n, b, l-1, usedMap)
	return res
}

func getPermutation(n int, k int) string {
	usedMap := make(map[int]bool)
	res := getPermutationImpl(n, k, n, usedMap)
	return res
}

// func main() {
// 	res := getPermutation(4,9)
// 	fmt.Println(res)
// }
// @lc code=end

