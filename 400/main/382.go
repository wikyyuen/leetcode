package main

import "math/rand"

//382. Linked List Random Node
//382. 链表随机节点

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

func (this *Solution) GetRandom() int {
	i, res := 0, 0
	p := this.head
	for p != nil {
		i++
		if rand.Intn(i) == 0 {
			res = p.Val
		}
		p = p.Next
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
