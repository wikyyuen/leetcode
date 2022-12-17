package main

import "math/rand"

//382. Linked List Random Node
//382. 链表随机节点

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	nodes []*ListNode
	size  int
}

func Constructor(head *ListNode) Solution {
	p := head
	solution := Solution{nodes: []*ListNode{}, size: 0}
	for p != nil {
		solution.size++
		solution.nodes = append(solution.nodes, p)
		p = p.Next
	}
	return solution
}

func (this *Solution) GetRandom() int {
	r := rand.Intn(this.size)
	return this.nodes[r].Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
