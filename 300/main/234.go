package main

//234. Palindrome Linked List
//234. 回文链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	left := head
	right := reverse(slow)
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true

}

func reverse(head *ListNode) *ListNode {
	var pre, next *ListNode
	pre = nil
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {

}
