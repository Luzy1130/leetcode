/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (46.05%)
 * Likes:    1663
 * Dislikes: 149
 * Total Accepted:    395K
 * Total Submissions: 827K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 * 
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preHead := &ListNode{ Val:0, Next:head }
	tmp := preHead
	pos1 := head
	pos2 := head.Next

	for pos1 != nil && pos2 != nil {
		pos1.Next = pos2.Next
		pos2.Next = pos1
		tmp.Next = pos2

		tmp = pos1
		pos1 = pos1.Next
		if pos1 != nil {
			pos2 = pos1.Next
		}
	}
	return preHead.Next
}
// @lc code=end

