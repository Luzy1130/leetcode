/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (36.65%)
 * Likes:    1305
 * Dislikes: 3084
 * Total Accepted:    419K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
 * one sorted array.
 *
 * Note:
 *
 *
 * The number of elements initialized in nums1 and nums2 are m and n
 * respectively.
 * You may assume that nums1 has enough space (size that is greater or equal to
 * m + n) to hold additional elements from nums2.
 *
 *
 * Example:
 *
 *
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * Output: [1,2,2,3,5,6]
 *
 *
 */
// package main

// import "fmt"

// func reverse(s []int) {
// 	if len(s) == 0 || len(s) == 1 {
// 		return
// 	}
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }
func merge(nums1 []int, m int, nums2 []int, n int) {
	// if n == 0 {
	// 	return
	// }
	// nums1 = append(nums1[:m], nums2...)
	// if m == 0 {
	// 	return
	// }

	// i := 0
	// j := m

	// for i != j && j < (m+n) {
	// 	if nums1[i] <= nums1[j] {
	// 		i++
	// 	} else {
	// 		reverse(nums1[i:j])
	// 		reverse(nums1[i : j+1])
	// 		j++
	// 	}
	// }

	i := m - 1
	j := n - 1
	k := n + m - 1

	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

// func main() {
// 	nums1 := []int{4, 0, 0, 0, 0, 0}
// 	nums2 := []int{1, 2, 3, 5, 6}

// 	// nums1 := []int{0}
// 	// nums2 := []int{1}

// 	merge(nums1, 1, nums2, 5)
// 	fmt.Println(nums1)
// }
