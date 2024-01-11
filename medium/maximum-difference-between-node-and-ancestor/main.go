package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxAncestorDiff(root *TreeNode) int {
	var res int

	var helper func(node *TreeNode, maxPrev int, minPrev int)
	helper = func(node *TreeNode, maxPrev int, minPrev int) {
		if node == nil {
			return
		}

		diffMax := int(math.Abs(float64(node.Val - maxPrev)))
		diffMin := int(math.Abs(float64(node.Val - minPrev)))

		maxPrev = max(maxPrev, node.Val)
		minPrev = min(minPrev, node.Val)

		res = max(res, diffMax)
		res = max(res, diffMin)

		helper(node.Left, maxPrev, minPrev)
		helper(node.Right, maxPrev, minPrev)
	}

	helper(root.Left, root.Val, root.Val)
	helper(root.Right, root.Val, root.Val)
	return res
}
