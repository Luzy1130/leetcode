/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (52.65%)
 * Likes:    3151
 * Dislikes: 222
 * Total Accepted:    553.5K
 * Total Submissions: 1M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Find the kth largest element in an unsorted array. Note that it is the kth
 * largest element in the sorted order, not the kth distinct element.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,2,1,5,6,4] and k = 2
 * Output: 5
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,2,3,1,2,4,5,5,6] and k = 4
 * Output: 4
 * 
 * Note: 
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 * 
 */
package main
import "fmt" 

// @lc code=start
// 弄个最大堆


func buildMaxHeap(nums[] int) []int {
	heap := make([]int, len(nums))

	for i, v := range nums {
		heap[i] = v
		pos := i + 1
		parent :=  pos / 2
		for parent >= 1 {
			if heap[pos-1] > heap[parent-1] {
				heap[pos-1], heap[parent-1] = heap[parent-1], heap[pos-1]
				pos = parent
				parent = pos / 2
			} else {
				break
			}
		}
	}

	return heap
}

func popMaxHeap(nums []int) (int, []int) {
	res := nums[0]
	l := len(nums)
	nums[0], nums[l-1] = nums[l-1], nums[0]
	nums = nums[0:(len(nums)-1)]
	l = len(nums)
	i := 1
	left := 2 * i
	right := 2 * i + 1
	for left <= l || right <= l {
		max := nums[left-1]
		j := left
		if right <= l && max < nums[right-1] {
			max = nums[right-1]
			j = right
		}

		if nums[i-1] >= max {
			break
		}
		
		nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
		i = j
		left = 2 * i
		right = 2 * i + 1
	}

	return res, nums
}

func findKthLargest(nums []int, k int) int {
	heap := buildMaxHeap(nums)
	// fmt.Println(heap)

	var res int
	for k > 0 {
		res, heap = popMaxHeap(heap)
		// fmt.Println(heap)
		k--
	}


	return res
}
// @lc code=end

func main() {
	res := findKthLargest([]int{-1, 2, 0}, 2)
	fmt.Println(res)
}