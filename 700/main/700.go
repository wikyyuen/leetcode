package main

//700. Search in a Binary Search Tree
//700. 二叉搜索树中的搜索

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}
