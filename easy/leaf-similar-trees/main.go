package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := make([]int, 0, 200)
	arr2 := make([]int, 0, 200)

	var helper func(node *TreeNode, first bool)

	helper = func(node *TreeNode, first bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if first {
				arr1 = append(arr1, node.Val)
			} else {
				arr2 = append(arr2, node.Val)
			}
			return
		}

		helper(node.Left, first)
		helper(node.Right, first)
	}

	helper(root1, true)
	helper(root2, false)

	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
