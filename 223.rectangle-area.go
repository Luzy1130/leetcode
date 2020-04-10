/*
 * @lc app=leetcode id=223 lang=golang
 *
 * [223] Rectangle Area
 *
 * https://leetcode.com/problems/rectangle-area/description/
 *
 * algorithms
 * Medium (37.13%)
 * Likes:    360
 * Dislikes: 661
 * Total Accepted:    101.7K
 * Total Submissions: 273K
 * Testcase Example:  '-3\n0\n3\n4\n0\n-1\n9\n2'
 *
 * Find the total area covered by two rectilinear rectangles in a 2D plane.
 * 
 * Each rectangle is defined by its bottom left corner and top right corner as
 * shown in the figure.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Input: A = -3, B = 0, C = 3, D = 4, E = 0, F = -1, G = 9, H = 2
 * Output: 45
 * 
 * Note:
 * 
 * Assume that the total area is never beyond the maximum possible value of
 * int.
 * 
 */
package main
import (
	"fmt"
	"sort"
)

// @lc code=start
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	// (A, B) (C, D) <---> (E, F) (G, H)

	area1 := (C - A) * (D - B)
	area2 := (G - E) * (H - F)
	// fmt.Println(area1, area2)
	// 先判断不重合的情况
	if C <= E || A >= G || B >= H || D <= F {
		return (area1 + area2)
	}

	var x []int
	x = append(x, A)
	x = append(x, C)
	x = append(x, E)
	x = append(x, G)
	sort.Sort(sort.IntSlice(x))
	

	var y []int
	y = append(y, B)
	y = append(y, D)
	y = append(y, F)
	y = append(y, H)
	sort.Sort(sort.IntSlice(y))


	res := area1 + area2 - (x[2]-x[1]) * (y[2] - y[1])
	return res
}
// @lc code=end

func main() {
	res := computeArea(-3, 0, 3, 4, 0, -1, 9, 2)
	fmt.Println(res)
}
