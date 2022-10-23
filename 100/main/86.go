package main

//86. Partition List
//分隔链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var frontFirstNode, frontLastNode, endFirstNode, endLastNode, midNode *ListNode

func partition(head *ListNode, x int) *ListNode {
	frontFirstNode = nil
	frontLastNode = nil
	endFirstNode = nil
	endLastNode = nil
	for head != nil {
		next := head.Next
		if head.Val < x {
			if frontFirstNode == nil {
				frontFirstNode = head
			}
			if frontLastNode == nil {
				frontLastNode = head
			} else {
				frontLastNode.Next = head
				frontLastNode = head
			}
		} else {
			if endFirstNode == nil {
				endFirstNode = head
			}
			if endLastNode == nil {
				endLastNode = head
			} else {
				endLastNode.Next = head
				endLastNode = head
			}
		}
		head.Next = nil
		head = next
	}
	if frontLastNode == nil {
		return endFirstNode
	}
	frontLastNode.Next = endFirstNode
	return frontFirstNode
}

func main() {

}
