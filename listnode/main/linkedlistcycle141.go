package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {
	list1 := &ListNode{Val: 1}
	list2 := &ListNode{Val: 2}
	list1.Next = list2

	result := hasCycle(list1)

	fmt.Println(result)
}
