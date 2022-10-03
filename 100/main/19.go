package main

//19. Remove Nth Node From End of List

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
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	i := 0
//	p1 := head
//	if p1 == nil {
//		return nil
//	}
//	p2 := head
//	for i <= n {
//		p2 = p2.Next
//		i++
//		if p2 == nil {
//			if i == n {
//				head = head.Next
//			} else {
//				head.Next = head.Next.Next
//			}
//
//			return head
//		}
//	}
//	for p2 != nil {
//		p2 = p2.Next
//		p1 = p1.Next
//	}
//	p1.Next = p1.Next.Next
//	return head
//}

//solution 2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	p2 := dummy
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return dummy.Next
}
func main() {

}
