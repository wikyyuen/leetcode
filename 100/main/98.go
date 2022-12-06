package main

//98. Validate Binary Search Tree
//98. 验证二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return checkNode(root, -2147483649, 2147483648)
}

func checkNode(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	if !checkNode(root.Left, min, root.Val) {
		return false
	}
	if !checkNode(root.Right, root.Val, max) {
		return false
	}
	return true
}
