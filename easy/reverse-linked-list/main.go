package main

import (
	"fmt"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

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

func ValListNodes(head *ListNode) []int {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func main() {
	var h, res *ListNode

	h = ConstructNodes([]int{1, 2, 3})
	res = reverseList(h)
	fmt.Println(ValListNodes(res), "[3 2 1]")
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		tmpNext := head.Next
		head.Next = prev
		prev = head
		head = tmpNext
	}

	return prev
}
