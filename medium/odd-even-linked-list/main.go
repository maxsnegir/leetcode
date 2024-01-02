package main

import (
	"fmt"

	ll "github.com/maxsnegir/leetcode/base"
)

func main() {
	res := oddEvenList(ll.ConstructNodes([]int{1, 2, 3, 4, 5}))
	fmt.Println(ll.ValListNodes(res), "[1,2,3,4,5]")

	res = oddEvenList(ll.ConstructNodes([]int{2, 1, 3, 5, 6, 4, 7}))
	fmt.Println(ll.ValListNodes(res), "[2,3,6,7,1,5,4]")
}

func oddEvenList(head *ll.ListNode) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmpHead := head
	odd := head
	tmpEven := head.Next
	even := head.Next

	for odd != nil && even != nil && odd.Next != nil && even.Next != nil {
		nextOdd := odd.Next.Next
		nextEven := even.Next.Next

		odd.Next = nextOdd
		even.Next = nextEven

		odd = nextOdd
		even = nextEven
	}

	odd.Next = tmpEven
	return tmpHead
}
