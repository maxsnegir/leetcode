package main

import (
	"fmt"

	ll "github.com/maxsnegir/leetcode/base"
)

func main() {
	res := ll.ConstructNodes([]int{5, 4, 2, 1})
	fmt.Println(pairSum(res), 6)

	res = ll.ConstructNodes([]int{4, 2, 2, 3})
	fmt.Println(pairSum(res), 7)
}

func pairSum(head *ll.ListNode) int {

	var max int
	left := head
	right := head

	n := 0
	cache := make(map[int]int)
	for right != nil && right.Next != nil {
		cache[n] = left.Val

		n++
		left = left.Next
		right = right.Next.Next
	}

	count := n * 2
	for left != nil {
		twinVal := cache[count-n-1]
		if twinVal+left.Val > max {
			max = twinVal + left.Val
		}
		left = left.Next
		n++
	}

	return max
}
