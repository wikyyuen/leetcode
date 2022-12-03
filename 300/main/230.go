package main

//230. Kth Smallest Element in a BST
//230. 二叉搜索树中第K小的元素

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var rank int

func kthSmallest(root *TreeNode, k int) int {
	res = 0
	rank = 0
	traverse(root, k)
	return res
}

func traverse(root *TreeNode, k int) {
	if root == nil || res != 0 {
		return
	}
	traverse(root.Left, k)
	rank++
	if rank == k {
		res = root.Val
		return
	}
	traverse(root.Right, k)
}
