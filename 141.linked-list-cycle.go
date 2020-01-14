/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 *
 * https://leetcode.com/problems/linked-list-cycle/description/
 *
 * algorithms
 * Easy (37.91%)
 * Likes:    2079
 * Dislikes: 307
 * Total Accepted:    507K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, determine if it has a cycle in it.
 * 
 * To represent a cycle in the given linked list, we use an integer pos which
 * represents the position (0-indexed) in the linked list where tail connects
 * to. If pos is -1, then there is no cycle in the linked list.
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [3,2,0,-4], pos = 1
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = [1,2], pos = 0
 * Output: true
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: head = [1], pos = -1
 * Output: false
 * Explanation: There is no cycle in the linked list.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Follow up:
 * 
 * Can you solve it using O(1) (i.e. constant) memory?
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
func hasCycle(head *ListNode) bool {
	// 使用两个指针，一个指针每次跳一步，一个指针每次跳两步
	// 如果两个指针能够在某一个时刻相等则是一个有圈的链表
	if head == nil || head.Next == nil {
		return false
	}
	// 两个指针起点相同
	pos1 := head
	pos2 := head
	for pos1 != nil && pos2 != nil {
		if pos1.Next != nil {
			pos1 = pos1.Next
		} else {
			break
		}

		if pos2.Next != nil && pos2.Next.Next != nil {
			pos2 = pos2.Next.Next
		} else {
			break
		}

		if pos1 == pos2 {
			return true
		}
	}

	return false

}
// @lc code=end

