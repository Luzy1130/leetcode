/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 * https://leetcode.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (56.45%)
 * Likes:    3244
 * Dislikes: 77
 * Total Accepted:    772.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Reverse a singly linked list.
 * 
 * Example:
 * 
 * 
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 * 
 * 
 * Follow up:
 * 
 * A linked list can be reversed either iteratively or recursively. Could you
 * implement both?
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

// 递归算法
func reverseListRecursive(head *ListNode) (newHeader *ListNode, newTail *ListNode) {
	if head == nil {
		return nil, nil
	}
	if head.Next == nil {
		return head, head
	}

	newHeader, newTail = reverseListRecursive(head.Next)
	head.Next = nil
	newTail.Next = head
	return newHeader, head
}

// 非递归算法
func reverseListNonRecursive(head *ListNode) (*ListNode) {
	if head == nil {
		return nil
	}
	curNode := head
	var newHeader *ListNode = nil

	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = newHeader
		newHeader = curNode
		curNode = nextNode
	}

	return newHeader
}

func reverseList(head *ListNode) *ListNode {
	// newHeader, _ := reverseListRecursive(head)

	newHeader := reverseListNonRecursive(head)


	return newHeader
}
// @lc code=end

