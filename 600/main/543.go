package main

//543. Diameter of Binary Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDeep int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDeep = 0
	traverse(root)
	return maxDeep
}

func getMaxDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := getMaxDeep(root.Left)
	right := getMaxDeep(root.Right)

	if left > right {
		return 1 + left
	} else {
		return 1 + right
	}
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	left := getMaxDeep(root.Left)
	right := getMaxDeep(root.Right)
	deep := left + right
	if maxDeep < deep {
		maxDeep = deep
	}
	traverse(root.Left)
	traverse(root.Right)
}
