/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 *
 * https://leetcode.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (39.79%)
 * Likes:    1855
 * Dislikes: 346
 * Total Accepted:    247.2K
 * Total Submissions: 613.1K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, reverse the nodes of a linked list k at a time and
 * return its modified list.
 * 
 * k is a positive integer and is less than or equal to the length of the
 * linked list. If the number of nodes is not a multiple of k then left-out
 * nodes in the end should remain as it is.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * Given this linked list: 1->2->3->4->5
 * 
 * For k = 2, you should return: 2->1->4->3->5
 * 
 * For k = 3, you should return: 3->2->1->4->5
 * 
 * Note:
 * 
 * 
 * Only constant extra memory is allowed.
 * You may not alter the values in the list's nodes, only nodes itself may be
 * changed.
 * 
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
// import "fmt"
func reverse(head *ListNode) *ListNode {
	guard := &ListNode{0, head}
	pos := head.Next
	head.Next = nil
	for pos != nil {
		next := pos.Next
		pos.Next = guard.Next
		guard.Next = pos
		pos = next
	}
	return guard.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 1 {
		return head
	}

	guard := &ListNode{0, head}
	cnt := 1
	preHead := guard
	pos := head
	for pos != nil && pos.Next != nil {
		for ;cnt < k && pos.Next != nil; cnt++ {
			pos = pos.Next
		}
		if cnt == k {
			// fmt.Println(pos)
			nextHead := pos.Next
			// reverse sub list
			pos.Next = nil
			newTail := preHead.Next
			newHead := reverse(preHead.Next)
			preHead.Next = newHead
			newTail.Next = nextHead
			preHead = newTail
			pos = nextHead
		}
		cnt = 1
	}

	return guard.Next
}
// @lc code=end

