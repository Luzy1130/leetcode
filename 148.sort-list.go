/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 *
 * https://leetcode.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (39.70%)
 * Likes:    2229
 * Dislikes: 113
 * Total Accepted:    238.2K
 * Total Submissions: 597.6K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Sort a linked list in O(n log n) time using constant space complexity.
 * 
 * Example 1:
 * 
 * 
 * Input: 4->2->1->3
 * Output: 1->2->3->4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -1->5->3->4->0
 * Output: -1->0->3->4->5
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 使用归并排序
func mergeSort(head *ListNode, l int) *ListNode {
	if l == 1 {
		return head
	}

	var list1 *ListNode = head
	var list2 *ListNode = nil

	pos := head
	cnt := 1
	for cnt < l / 2 {
		pos = pos.Next
		cnt++
	}
	list2 = pos.Next
	pos.Next = nil

	afterSort1 := mergeSort(list1, cnt)
	afterSort2 := mergeSort(list2, l-cnt)

	// 合并
	var newPreHead *ListNode = &ListNode{}
	pos = newPreHead
	for afterSort1 != nil && afterSort2 != nil {
		var tmp *ListNode = nil
 		if afterSort1.Val < afterSort2.Val {
			tmp = afterSort1
			afterSort1 = afterSort1.Next
		} else {
			tmp = afterSort2
			afterSort2 = afterSort2.Next
		}
		
		pos.Next = tmp
		pos = pos.Next
	}

	if afterSort1 != nil {
		pos.Next = afterSort1
	}
	if afterSort2 != nil {
		pos.Next = afterSort2
	}
	return newPreHead.Next
}

func sortList(head *ListNode) *ListNode {
	l := 0
	for pos := head; pos != nil; pos = pos.Next {
		l++
	}
	if l == 0 {
		return head
	}

	return mergeSort(head, l)
}
// @lc code=end

