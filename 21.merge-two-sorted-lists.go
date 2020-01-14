/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (48.58%)
 * Likes:    2575
 * Dislikes: 372
 * Total Accepted:    665.3K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 * 
 * Example:
 * 
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode
	var pos *ListNode = &head

	node1 := l1
	node2 := l2
	for ; node1 != nil && node2 != nil; {
		if node1.Val <= node2.Val {
			pos.Next = node1
			node1 = node1.Next
		} else {
			pos.Next = node2
			node2 = node2.Next
		}
		
		pos = pos.Next
	}

	if node1 != nil {
		pos.Next = node1
	}
	if node2 != nil {
		pos.Next = node2
	}

	return head.Next
}

