/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 *
 * https://leetcode.com/problems/implement-stack-using-queues/description/
 *
 * algorithms
 * Easy (40.52%)
 * Likes:    454
 * Dislikes: 495
 * Total Accepted:    153.9K
 * Total Submissions: 365.3K
 * Testcase Example:  '["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a stack using queues.
 * 
 * 
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * empty() -- Return whether the stack is empty.
 * 
 * 
 * Example:
 * 
 * 
 * MyStack stack = new MyStack();
 * 
 * stack.push(1);
 * stack.push(2);  
 * stack.top();   // returns 2
 * stack.pop();   // returns 2
 * stack.empty(); // returns false
 * 
 * Notes:
 * 
 * 
 * You must use only standard operations of a queue -- which means only push to
 * back, peek/pop from front, size, and is empty operations are valid.
 * Depending on your language, queue may not be supported natively. You may
 * simulate a queue by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a queue.
 * You may assume that all operations are valid (for example, no pop or top
 * operations will be called on an empty stack).
 * 
 * 
 */

// @lc code=start
type MyStack struct {
	queues 		[][]int
	len	  		int
	topIndex 	int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	stack := MyStack{}
	stack.queues = append(stack.queues, []int{})
	stack.queues = append(stack.queues, []int{})
	stack.len = 0
	stack.topIndex = 0

	return stack
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	topIndex := this.topIndex
	this.queues[topIndex] = append(this.queues[topIndex], x)
	if len(this.queues[topIndex]) > 1 {
		index := 1 ^ topIndex
		this.queues[index] = append(this.queues[index], this.queues[topIndex][0])
		this.queues[topIndex] = this.queues[topIndex][1:len(this.queues[topIndex])]
	}
	this.len++
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	topIndex := this.topIndex
	index := 1 ^ topIndex

	res := this.queues[topIndex][0]
	this.queues[topIndex] = this.queues[topIndex][0:0]
	for i := 0; i < len(this.queues[index])-1; i++ {
		this.queues[topIndex] = append(this.queues[topIndex], this.queues[index][i])
	}

	if len(this.queues[index]) > 0 {
		this.queues[index] = this.queues[index][len(this.queues[index])-1:]
	}

	this.topIndex = index
	this.len--
	return res
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.queues[this.topIndex][0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    // for i := 0; len(this.queues); i++ {
	// 	if len(this.queues[i]) != 0 {
	// 		return false
	// 	}
	// }
	// return true
	return this.len == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

