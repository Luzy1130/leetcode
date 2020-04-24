/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 *
 * https://leetcode.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (52.92%)
 * Likes:    1205
 * Dislikes: 143
 * Total Accepted:    141K
 * Total Submissions: 264.7K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The most significant digit comes first and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 * 
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 * 
 * Follow up:
 * What if you cannot modify the input lists? In other words, reversing the
 * lists is not allowed.
 * 
 * 
 * 
 * Example:
 * 
 * Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 8 -> 0 -> 7
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
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 不能将原链表翻转
    // 用两个数组来辅助
    var num1 []int
    var num2 []int
    for pos := l1; pos != nil; pos = pos.Next {
        num1 = append(num1, pos.Val)
    }
    for pos := l2; pos != nil; pos = pos.Next {
        num2 = append(num2, pos.Val)
    }
    
    var newHead *ListNode = &ListNode{0, nil}
    i := len(num1)-1
    j := len(num2)-1
    carry := 0
    for i >= 0 && j >= 0 {
        sum := carry + num1[i] + num2[j]
        carry = sum / 10
        sum = sum % 10
        tmp := &ListNode{sum, newHead.Next}
        newHead.Next = tmp
        i--
        j--
    }
    
    k := -1
    var num []int
    if i >= 0 {
        k = i
        num = num1
    }
    if j >= 0 {
        k = j
        num = num2
    }
    
    for k >= 0 {
        sum := carry + num[k]
        carry = sum / 10
        sum = sum % 10
        tmp := &ListNode{sum, newHead.Next}
        newHead.Next = tmp
        k--
    }
    
    if carry != 0 {
        tmp := &ListNode{carry, newHead.Next}
        newHead.Next = tmp
    }
    return newHead.Next
}
// @lc code=end

