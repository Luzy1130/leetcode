/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 *
 * https://leetcode.com/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (44.60%)
 * Likes:    761
 * Dislikes: 126
 * Total Accepted:    180.8K
 * Total Submissions: 390.4K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a queue using stacks.
 * 
 * 
 * push(x) -- Push element x to the back of queue.
 * pop() -- Removes the element from in front of queue.
 * peek() -- Get the front element.
 * empty() -- Return whether the queue is empty.
 * 
 * 
 * Example:
 * 
 * 
 * MyQueue queue = new MyQueue();
 * 
 * queue.push(1);
 * queue.push(2);  
 * queue.peek();  // returns 1
 * queue.pop();   // returns 1
 * queue.empty(); // returns false
 * 
 * Notes:
 * 
 * 
 * You must use only standard operations of a stack -- which means only push to
 * top, peek/pop from top, size, and is empty operations are valid.
 * Depending on your language, stack may not be supported natively. You may
 * simulate a stack by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a stack.
 * You may assume that all operations are valid (for example, no pop or peek
 * operations will be called on an empty queue).
 * 
 * 
 */

// @lc code=start
type MyQueue struct {
	stacks [][]int
	peekIndex int
	len int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	queue := MyQueue{}
	queue.stacks = append(queue.stacks, []int{})
	queue.stacks = append(queue.stacks, []int{})
	queue.peekIndex = 0
	queue.len = 0
	
	return queue
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	peekIndex := this.peekIndex
	otherIndex := 1 ^ peekIndex

	this.stacks[peekIndex] = append(this.stacks[peekIndex], x)
	if len(this.stacks[peekIndex]) > 1 {
		num := this.stacks[peekIndex][1]
		this.stacks[otherIndex] = append(this.stacks[otherIndex], num)
		this.stacks[peekIndex] = this.stacks[peekIndex][0:1]
	}

	this.len++
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	peekIndex := this.peekIndex
	otherIndex := 1 ^ peekIndex
	res := this.stacks[peekIndex][0]

	this.stacks[peekIndex] = this.stacks[peekIndex][0:0]
	for i := len(this.stacks[otherIndex])-1; i > 0; i-- {
		num := this.stacks[otherIndex][i]
		this.stacks[peekIndex] = append(this.stacks[peekIndex], num)
	}

	if len(this.stacks[otherIndex]) > 0 {
		this.stacks[otherIndex] = this.stacks[otherIndex][0:1]
	}

	this.peekIndex = otherIndex

	this.len--
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	peekIndex := this.peekIndex
	return this.stacks[peekIndex][0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.len == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

