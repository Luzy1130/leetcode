/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (35.74%)
 * Likes:    3820
 * Dislikes: 248
 * Total Accepted:    556.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * Merge k sorted linked lists and return it as one sorted list. Analyze and
 * describe its complexity.
 * 
 * Example:
 * 
 * 
 * Input:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
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
func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	pos := preHead
	pos1 := list1
	pos2 := list2

	for pos1 != nil && pos2 != nil {
		var tmp *ListNode
		if pos1.Val <= pos2.Val {
			tmp = pos1
			pos1 = pos1.Next
		} else {
			tmp = pos2
			pos2 = pos2.Next
		}
		pos.Next = tmp
		pos = pos.Next
		pos.Next = nil
	}

	if pos1 != nil {
		pos.Next = pos1
	}
	if pos2 != nil {
		pos.Next = pos2
	}

	return preHead.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	// 两条链表直接合并
	if len(lists) == 2 {
		res := merge2Lists(lists[0], lists[1])
		return res
	}
	n := len(lists)
	list1 := mergeKLists(lists[0:((n+1)/2)])
	list2 := mergeKLists(lists[((n+1)/2):])

	res := merge2Lists(list1, list2)
	return res
}
// @lc code=end

