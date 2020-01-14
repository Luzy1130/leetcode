/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	var valueSet = make(map[int]int)
	var ret []int
	for i, num := range nums {   
		if j, ok := valueSet[target-num]; ok {
				return []int{j, i}
		}
		valueSet[num] = i
	}
	
	return ret
}

