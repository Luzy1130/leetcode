/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.54%)
 * Likes:    2493
 * Dislikes: 184
 * Total Accepted:    506.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 * 
 * Example:
 * 
 * 
 * Given linked list: 1->2->3->4->5, and n = 2.
 * 
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 * 
 * 
 * Note:
 * 
 * Given n will always be valid.
 * 
 * Follow up:
 * 
 * Could you do this in one pass?
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
 // 两个指针互追
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// head之前插入一个伪队列头，便于操作
	tmpHead := ListNode{Val:0, Next:head}
	pos1 := &tmpHead
	pos2 := &tmpHead

	for i := 0; i < n; i++ {
		pos1 = pos1.Next
	}
	for pos1.Next != nil {
		pos1 = pos1.Next
		pos2 = pos2.Next
	}

	pos2.Next = pos2.Next.Next
	return tmpHead.Next
}
// @lc code=end

