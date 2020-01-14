/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (38.43%)
 * Likes:    2414
 * Dislikes: 253
 * Total Accepted:    387.6K
 * Total Submissions: 965.7K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 * 
 * 
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * 
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 * 
 * 
 * 
 * 
 */

// @lc code=start
type MinStack struct {
	data     []int
	minimals []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	stack := new(MinStack)
	return *stack
}


func (this *MinStack) Push(x int)  {
	this.data = append(this.data, x)
	if len(this.minimals) == 0 || x <= this.GetMin() {
		this.minimals = append(this.minimals, x)
	}
}


func (this *MinStack) Pop()  {
	x := this.data[len(this.data)-1]
	this.data = this.data[0:(len(this.data)-1)]
	if x == this.GetMin() {
		this.minimals = this.minimals[0:(len(this.minimals)-1)]
	}
}


func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return -1
}


func (this *MinStack) GetMin() int {
	if len(this.minimals) > 0 {
		return this.minimals[len(this.minimals)-1]
	}
	return this.Top()
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

