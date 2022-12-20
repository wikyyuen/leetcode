package main

//1038. Binary Search Tree to Greater Sum Tree
//1038. 从二叉搜索树到更大和树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var sum int

func bstToGst(root *TreeNode) *TreeNode {
	sum = 0
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Right)
	sum += root.Val
	root.Val = sum
	traverse(root.Left)
}
