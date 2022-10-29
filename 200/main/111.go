package main

import "fmt"

//111. Minimum Depth of Binary Tree
//111. 二叉树的最小深度

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	curNodeList := []*TreeNode{}
	curNodeList = append(curNodeList, root)
	step := 1
	for len(curNodeList) > 0 {
		sz := len(curNodeList)
		for i := 0; i < sz; i++ {
			node := curNodeList[0]
			curNodeList = curNodeList[1:]
			if node.Left == nil && node.Right == nil {
				return step
			}
			if node.Left != nil {
				curNodeList = append(curNodeList, node.Left)
			}
			if node.Right != nil {
				curNodeList = append(curNodeList, node.Right)
			}
		}
		step++
	}
	return step
}

func main() {
	tree1, tree2, tree3, tree4, tree5 := &TreeNode{}, &TreeNode{}, &TreeNode{}, &TreeNode{}, &TreeNode{}
	tree1.Val = 1
	tree2.Val = 2
	tree3.Val = 3
	tree4.Val = 4
	tree5.Val = 5
	tree1.Left = tree2
	tree1.Right = tree3
	tree2.Left = tree4
	tree2.Right = tree5

	result := minDepth(tree1)

	fmt.Println(result)
}
