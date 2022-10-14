package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func preorderTraversal(root *TreeNode) []int {
	res = []int{}
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	traverse(root.Left)
	traverse(root.Right)
}
