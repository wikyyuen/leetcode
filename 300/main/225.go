package main

//225. Implement Stack using Queues
//225. 用队列实现栈

type MyStack struct {
	items []int
	lens  int
}

func Constructor() MyStack {
	stack := new(MyStack)
	stack.lens = 0
	stack.items = make([]int, 0)
	return *stack
}

func (this *MyStack) Push(x int) {
	this.items = append(this.items, x)
	this.lens += 1
}

func (this *MyStack) Pop() int {
	x := this.items[this.lens-1]
	this.items = this.items[:this.lens-1]
	this.lens -= 1
	return x
}

func (this *MyStack) Top() int {
	return this.items[this.lens-1]
}

func (this *MyStack) Empty() bool {
	return this.lens == 0
}
