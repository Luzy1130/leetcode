/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 *
 * https://leetcode.com/problems/odd-even-linked-list/description/
 *
 * algorithms
 * Medium (52.29%)
 * Likes:    1382
 * Dislikes: 275
 * Total Accepted:    214.2K
 * Total Submissions: 405.4K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given a singly linked list, group all odd nodes together followed by the
 * even nodes. Please note here we are talking about the node number and not
 * the value in the nodes.
 * 
 * You should try to do it in place. The program should run in O(1) space
 * complexity and O(nodes) time complexity.
 * 
 * Example 1:
 * 
 * 
 * Input: 1->2->3->4->5->NULL
 * Output: 1->3->5->2->4->NULL
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 2->1->3->5->6->4->7->NULL
 * Output: 2->3->6->7->1->5->4->NULL
 * 
 * 
 * Note:
 * 
 * 
 * The relative order inside both the even and odd groups should remain as it
 * was in the input.
 * The first node is considered odd, the second node even and so on ...
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
 func oddEvenList(head *ListNode) *ListNode {
    // 原来的列表顺序的奇偶来归纳
    if head == nil || head.Next == nil {
        return head
    }
    guard := &ListNode{0, head}
    cnt := 3 // 用于记录当前访问的位置index
    pos := head.Next.Next // 用于记录当前访问的node
    odd := head // 用于记录odd链表的最后一个node
    even := head.Next // 用于记录even链表的最后一个node
    
    for pos != nil {
        if cnt % 2 == 1 {
            next := pos.Next
            pos.Next = odd.Next
            odd.Next = pos
            odd = odd.Next
            even.Next = next
            even = even.Next
            pos = next
        } else {
            pos = pos.Next
        }
        cnt++
    }
    
    return guard.Next
}
// @lc code=end

