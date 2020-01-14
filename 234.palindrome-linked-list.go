/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (36.86%)
 * Likes:    2245
 * Dislikes: 308
 * Total Accepted:    333.4K
 * Total Submissions: 886.1K
 * Testcase Example:  '[1,2]'
 *
 * Given a singly linked list, determine if it is a palindrome.
 * 
 * Example 1:
 * 
 * 
 * Input: 1->2
 * Output: false
 * 
 * Example 2:
 * 
 * 
 * Input: 1->2->2->1
 * Output: true
 * 
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
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
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return true
	}

	mid := head
	fast := head
	for fast != nil && fast.Next != nil {
		// 节点数为奇数时mid为中间节点，节点为偶数时mid为中间偏后一个节点
		mid = mid.Next
		fast = fast.Next.Next
	}

	fast = head.Next
	newHead := head
	nextNode := newHead.Next
	newHead.Next = nil
	// 翻转mid之前的链表
	for nextNode != mid && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		
		curNode := newHead
		newHead = nextNode
		nextNode = nextNode.Next
		newHead.Next = curNode
	}
	if fast != nil && fast.Next != nil {
		mid = mid.Next
	}

	for newHead != nil && mid != nil {
		if newHead.Val != mid.Val {
			return false
		}

		newHead = newHead.Next
		mid = mid.Next
	}
	return true
}
// @lc code=end

