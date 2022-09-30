package main

import "fmt"

//239. Sliding Window Maximum

type doubleListNode struct {
	Prev *doubleListNode
	Next *doubleListNode
	Val  int
}

type Window struct {
	KeyMap    []*doubleListNode
	FirstNode *doubleListNode
	LastNode  *doubleListNode
	Capacity  int
}

func (w *Window) Pull(a int) {
	length := len(w.KeyMap)
	if length == w.Capacity {
		listNode := w.KeyMap[0]
		w.KeyMap = w.KeyMap[1:]
		if listNode.Next != nil {
			listNode.Next.Prev = listNode.Prev
		}
		if listNode.Prev != nil {
			listNode.Prev.Next = listNode.Next
		}
		if w.FirstNode == listNode {
			w.FirstNode = listNode.Next
		}
		if w.LastNode == listNode {
			w.LastNode = listNode.Prev
		}
	}
	listNode := &doubleListNode{
		Val: a,
	}
	if w.LastNode == nil {
		w.LastNode = listNode
		w.FirstNode = listNode
	} else {
		prevNode := w.LastNode
		for prevNode != nil {
			if prevNode.Val > a {
				prevNode.Next = listNode
				listNode.Prev = prevNode
				break
			} else {
				if prevNode.Prev != nil {
					prevNode.Next = nil
					prevNode = prevNode.Prev
					prevNode.Next.Prev = nil
					prevNode.Next = nil
				} else {
					prevNode.Prev = nil
					prevNode.Next = nil
					prevNode = listNode
					break
				}
			}
		}
		w.LastNode = listNode
	}
	if w.FirstNode == nil || w.FirstNode.Val <= a {
		w.FirstNode = listNode
	}
	w.KeyMap = append(w.KeyMap, listNode)
}

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	window := &Window{
		KeyMap:   []*doubleListNode{},
		Capacity: k,
	}
	for _, i2 := range nums {
		window.Pull(i2)
		if len(window.KeyMap) == k {
			res = append(res, window.FirstNode.Val)
		}
	}
	return res
}

func main() {
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//k := 3
	nums := []int{5, 3, 4}
	k := 1
	result := maxSlidingWindow(nums, k)
	fmt.Println(result)
}
