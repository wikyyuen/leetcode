package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	change(root)
}

func change(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftLast := change(root.Left)
	last := change(root.Right)
	if root.Left == nil && root.Right == nil {
		return root
	}
	if root.Left != nil {
		if root.Right != nil {
			leftLast.Right = root.Right
		}
		root.Right = root.Left
		root.Left = nil
	}
	if last == nil {
		last = leftLast
	}
	return last
}

func main() {

}
