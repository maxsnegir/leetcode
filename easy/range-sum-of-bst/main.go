package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var res int

	var helper func(node *TreeNode)

	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		helper(node.Left)
		helper(node.Right)

		if node.Val >= low && node.Val <= high {
			res += node.Val
		}
	}

	helper(root)
	return res
}
