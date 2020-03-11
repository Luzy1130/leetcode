/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 *
 * https://leetcode.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (38.23%)
 * Likes:    1010
 * Dislikes: 260
 * Total Accepted:    196.8K
 * Total Submissions: 493.6K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * Given a linked list and a value x, partition it such that all nodes less
 * than x come before nodes greater than or equal to x.
 * 
 * You should preserve the original relative order of the nodes in each of the
 * two partitions.
 * 
 * Example:
 * 
 * 
 * Input: head = 1->4->3->2->5->2, x = 3
 * Output: 1->2->2->4->3->5
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

 // 题目的意思是：将所有小于x的节点放到大于等于x的节点之前
 // 根据大小，将元素先放到两个链表中，最后合并链表
func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{0, nil}
	afterHead := &ListNode{0, nil}
	

	beforePre := beforeHead
	afterPre := afterHead
	for pos := head; pos != nil; pos = pos.Next {
		if pos.Val < x {
			beforePre.Next = pos
			beforePre = pos
		} else {
			afterPre.Next = pos
			afterPre = pos
		}
	}

	afterPre.Next = nil
	beforePre.Next = afterHead.Next
	return beforeHead.Next
}
// @lc code=end

