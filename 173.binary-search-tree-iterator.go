/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 *
 * https://leetcode.com/problems/binary-search-tree-iterator/description/
 *
 * algorithms
 * Medium (53.46%)
 * Likes:    2092
 * Dislikes: 265
 * Total Accepted:    273.1K
 * Total Submissions: 508.2K
 * Testcase Example:  '["BSTIterator","next","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]\n[[[7,3,15,null,null,9,20]],[null],[null],[null],[null],[null],[null],[null],[null],[null]]'
 *
 * Implement an iterator over a binary search tree (BST). Your iterator will be
 * initialized with the root node of a BST.
 * 
 * Calling next() will return the next smallest number in the BST.
 * 
 * 
 * 
 * 
 * 
 * 
 * Example:
 * 
 * 
 * 
 * 
 * BSTIterator iterator = new BSTIterator(root);
 * iterator.next();    // return 3
 * iterator.next();    // return 7
 * iterator.hasNext(); // return true
 * iterator.next();    // return 9
 * iterator.hasNext(); // return true
 * iterator.next();    // return 15
 * iterator.hasNext(); // return true
 * iterator.next();    // return 20
 * iterator.hasNext(); // return false
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * next() and hasNext() should run in average O(1) time and uses O(h) memory,
 * where h is the height of the tree.
 * You may assume that next() call will always be valid, that is, there will be
 * at least a next smallest number in the BST when next() is called.
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	Root *TreeNode
	Pos  *TreeNode
    Stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	t := BSTIterator{}
	t.Root = root
	pos := root
	stack := []*TreeNode{}

	for pos != nil {
		stack = append(stack, pos)
		pos = pos.Left
	}

	if len(stack) != 0 {
		pos = stack[len(stack)-1].Right
	}

	t.Pos = pos
	t.Stack = stack

	return t
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	stack := this.Stack
	node := stack[len(stack)-1]
	stack = stack[0:(len(stack)-1)]
	pos := this.Pos

	if pos != nil {
		for pos != nil {
			stack = append(stack, pos)
			pos = pos.Left
		}
	}

	if len(stack) != 0 {
		pos = stack[len(stack)-1].Right
	}
	
	this.Pos = pos	
	this.Stack = stack

	// fmt.Println(node.Val)
	return node.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.Stack) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

