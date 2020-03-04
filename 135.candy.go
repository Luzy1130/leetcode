/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 *
 * https://leetcode.com/problems/candy/description/
 *
 * algorithms
 * Hard (29.12%)
 * Likes:    732
 * Dislikes: 147
 * Total Accepted:    120.2K
 * Total Submissions: 396.3K
 * Testcase Example:  '[1,0,2]'
 *
 * There are N children standing in a line. Each child is assigned a rating
 * value.
 * 
 * You are giving candies to these children subjected to the following
 * requirements:
 * 
 * 
 * Each child must have at least one candy.
 * Children with a higher rating get more candies than their neighbors.
 * 
 * 
 * What is the minimum candies you must give?
 * 
 * Example 1:
 * 
 * 
 * Input: [1,0,2]
 * Output: 5
 * Explanation: You can allocate to the first, second and third child with 2,
 * 1, 2 candies respectively.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [1,2,2]
 * Output: 4
 * Explanation: You can allocate to the first, second and third child with 1,
 * 2, 1 candies respectively.
 * ⁠            The third child gets 1 candy because it satisfies the above two
 * conditions.
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	var output int = 0
	down := make([]int, len(ratings))
	candies := make([]int, len(ratings))
	for i := len(ratings)-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			down[i] = down[i+1] + 1
		}
	}

	for i := 0; i < len(ratings); i++ {
		if i == 0 {
			candies[i] = 1 + down[i]
		} else {
			var tmp int
			if ratings[i] == ratings[i-1] {
				tmp = 1
			} else if ratings[i] > ratings[i-1] {
				if down[i] == 0 {
					tmp = candies[i-1] + 1
				} else {
					if candies[i-1] >= down[i] {
						tmp = 1 + candies[i-1] - down[i]
					} else {
						tmp = 1
					}
				}
			} else {
				tmp = 1
			}
			candies[i] = tmp + down[i]
		}

		output += candies[i]
	}
	// fmt.Println(down)
	// fmt.Println(candies)
	return output
}

// func main() {
// 	output := candy([]int{4,2,1,4,10,12,4,3,6,7,1})	// 24
// //  output := candy([]int{1,6,10,8,7,3,2}) //18
// 	fmt.Println(output)
// }
// @lc code=end

