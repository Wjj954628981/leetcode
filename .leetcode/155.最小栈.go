/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (49.03%)
 * Likes:    261
 * Dislikes: 0
 * Total Accepted:    39K
 * Total Submissions: 78.7K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 *
 * push(x) -- 将元素 x 推入栈中。
 * pop() -- 删除栈顶的元素。
 * top() -- 获取栈顶元素。
 * getMin() -- 检索栈中的最小元素。
 *
 *
 * 示例:
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
 */
type MinStack struct {
	Data   []int
	Helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Data:   []int{},
		Helper: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	if len(this.Helper) == 0 || x <= this.Helper[len(this.Helper)-1] {
		this.Helper = append(this.Helper, x)
	}
}

func (this *MinStack) Pop() {
	top := this.Data[len(this.Data)-1]
	if top == this.Helper[len(this.Helper)-1] {
		this.Helper = this.Helper[:len(this.Helper)-1]
	}
	this.Data = this.Data[:len(this.Data)-1]
}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	return this.Helper[len(this.Helper)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

