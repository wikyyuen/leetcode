package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//solution1
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	next := head.Next
//	result := reverseList(next)
//	next.Next = head
//	head.Next = nil
//	return result
//}

func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next

	}
	return pre
}
func main() {
	list1 := &ListNode{
		Val: 1,
	}
	list2 := &ListNode{
		Val: 2,
	}
	list1.Next = list2
	result := reverseList(list1)
	fmt.Println(result)
}
