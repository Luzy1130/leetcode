/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
    node := head
    carry := 0
    node1 := l1
    node2 := l2
    
    for ; node1 != nil && node2 != nil; {
        sum := node1.Val + node2.Val + carry
        node.Val = sum % 10
        carry = sum / 10
        if carry != 0 || node1.Next != nil || node2.Next != nil {
            node.Next = new(ListNode)
            node = node.Next
        }
        node1=node1.Next
        node2=node2.Next
    }
    
    var nodeTmp *ListNode
    if node1 != nil {
        nodeTmp = node1
    } else if node2 != nil {
        nodeTmp = node2
    }
    for ; nodeTmp != nil; nodeTmp = nodeTmp.Next {
        sum := nodeTmp.Val + carry
        node.Val = sum % 10
        carry = sum / 10
        if carry != 0 || nodeTmp.Next != nil {
            node.Next = new(ListNode)
            node = node.Next
        }
    }
    
    if carry != 0 {
        node.Val = carry
    }
    
    return head
}

