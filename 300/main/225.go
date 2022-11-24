package main

//225. Implement Stack using Queues
//225. 用队列实现栈

type MyStack struct {
	items []int
}

func Constructor() MyStack {
	return MyStack{items: []int{}}
}

func (this *MyStack) Push(x int) {
	this.items = append(this.items, x)
}

func (this *MyStack) Pop() int {
	x := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	return x
}

func (this *MyStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.items) == 0
}
