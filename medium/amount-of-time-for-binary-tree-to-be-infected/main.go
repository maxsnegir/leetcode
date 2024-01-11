package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}

	var helper func(node *TreeNode, path int, prev int) int
	helper = func(node *TreeNode, path int, prev int) int {
		var leftPath, rightPath int
		if node == nil {
			return path
		}
		if prev == start {
			path--
		}

		if node.Left != nil {
			leftPath = helper(node.Left, path+1, node.Val)
		}
		if node.Right != nil {
			rightPath = helper(node.Right, path+1, node.Val)
		}

		return max(leftPath, rightPath)
	}

	return helper(root, 0, 0)
}
