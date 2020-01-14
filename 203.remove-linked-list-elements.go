/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (36.32%)
 * Likes:    1115
 * Dislikes: 61
 * Total Accepted:    270.8K
 * Total Submissions: 735.1K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 * 
 * Example:
 * 
 * 
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
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
func removeElements(head *ListNode, val int) *ListNode {
	preHead := &ListNode{}
	preHead.Next = head

	preNode := preHead
	curNode := preNode.Next

	for curNode != nil {
		if curNode.Val == val {
			preNode.Next = curNode.Next
		} else {
			preNode = preNode.Next
		}
		curNode = curNode.Next
	}

	return preHead.Next
}
// @lc code=end

