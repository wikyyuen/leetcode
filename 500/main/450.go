package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		//删除操作
		//和右节点的最小互换
		if root.Right == nil {
			root = root.Left
		} else {
			p := root.Right
			if p.Left == nil {
				p.Left = root.Left
				root = p
			} else {
				for p.Left.Left != nil {
					p = p.Left
				}
				p2 := p.Left
				p.Left = p2.Right
				p2.Right = root.Right
				p2.Left = root.Left
				root = p2
				//root.Val = p.Left.Val
				//p.Left = p.Left.Right
			}
		}
	} else if root.Val > key {
		//BST
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func main() {

}
