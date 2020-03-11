/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (33.90%)
 * Likes:    1298
 * Dislikes: 98
 * Total Accepted:    226K
 * Total Submissions: 638.3K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 * 
 * Return the linked list sorted as well.
 * 
 * Example 1:
 * 
 * 
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 1->1->1->2->3
 * Output: 2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	preHead := &ListNode{0, head}
	pre := preHead
	pos1 := head
	pos2 := pos1.Next
	waitingDelete := false

	for pos1 != nil {
		if pos2 != nil && pos1.Val == pos2.Val {
			pos2 = pos2.Next
			waitingDelete = true
		} else {
			if waitingDelete {
				pre.Next = pos2
				pos1 = pos2
				if pos2 != nil {
					pos2 = pos2.Next
				}
				waitingDelete = false
			} else {
				pre = pos1
				pos1 = pos2
				if pos2 != nil {
					pos2 = pos2.Next
				}
			}
		}
	}

	return preHead.Next
}
// @lc code=end

