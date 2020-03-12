/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (35.85%)
 * Likes:    1860
 * Dislikes: 121
 * Total Accepted:    244.6K
 * Total Submissions: 654.5K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Reverse a linked list from position m to n. Do it in one-pass.
 * 
 * Note: 1 ≤ m ≤ n ≤ length of list.
 * 
 * Example:
 * 
 * 
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	i := 1
	preHead := &ListNode{0, head}
	pre := preHead
	pos := head

	for ; i < m; i++ {
		pre = pos
		pos = pos.Next
	}

	tail := pos
	curHead := pos
	tmp := curHead
	next := curHead.Next
	// i = 2
	for ; i < n; i++ {
		curHead = next
		next = next.Next
		curHead.Next = tmp
		tmp = curHead
	}
	tail.Next= next
	pre.Next = curHead

	return preHead.Next
}
// @lc code=end

