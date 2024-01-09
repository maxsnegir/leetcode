package main

import (
	"fmt"
)

func ConstructNodes(values []int) *ListNode {
	tmpHead := &ListNode{}
	dummy := &ListNode{Next: tmpHead}

	for _, val := range values {
		next := &ListNode{Val: val}
		tmpHead.Next = next
		tmpHead = next
	}
	return dummy.Next.Next
}

type ListNode struct {
	Next *ListNode
	Val  int
}

func ValListNodes(head *ListNode) []int {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func main() {
	h1 := ConstructNodes([]int{1, 4, 5})
	h2 := ConstructNodes([]int{1, 3, 4})
	h3 := ConstructNodes([]int{2, 6})

	res := mergeKLists([]*ListNode{h1, h2, h3})
	fmt.Println(ValListNodes(res), "[1 1 2 3 4 4 5 6]")

	res = mergeKLists([]*ListNode{})
	fmt.Println(ValListNodes(res), "[]")

	res = mergeKLists([]*ListNode{nil})
	fmt.Println(ValListNodes(res), "[]")

}

func mergeKLists(lists []*ListNode) *ListNode {
	heads := make([]*ListNode, 0, len(lists))
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}

		heads = append(heads, lists[i])
	}

	next := &ListNode{}
	dummy := &ListNode{Next: next}

	for {
		idx := 0
		var minimum *ListNode

		for i, head := range heads {
			if head == nil {
				continue
			}
			if minimum == nil || head.Val < minimum.Val {
				minimum = head
				idx = i
			}
		}

		if minimum == nil {
			break
		}

		next.Next = minimum
		next = next.Next

		if heads[idx] != nil {
			heads[idx] = heads[idx].Next
		}
	}

	return dummy.Next.Next
}
