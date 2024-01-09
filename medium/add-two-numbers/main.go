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
	var h1, h2, res *ListNode

	h1 = ConstructNodes([]int{2, 4, 3})
	h2 = ConstructNodes([]int{5, 6, 4})
	res = addTwoNumbers(h1, h2)
	fmt.Println(ValListNodes(res), "[7 0 8]")

	h1 = ConstructNodes([]int{9, 9, 9, 9, 9, 9, 9})
	h2 = ConstructNodes([]int{9, 9, 9, 9})
	res = addTwoNumbers(h1, h2)
	fmt.Println(ValListNodes(res), "[8 9 9 9 0 0 0 1]")

	h1 = ConstructNodes([]int{0})
	h2 = ConstructNodes([]int{0})
	res = addTwoNumbers(h1, h2)
	fmt.Println(ValListNodes(res), "[0]")

	h1 = ConstructNodes([]int{2, 4, 9})
	h2 = ConstructNodes([]int{5, 6, 4, 9})
	res = addTwoNumbers(h1, h2)
	fmt.Println(ValListNodes(res), "[7 0 4 0 1]")

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	next := &ListNode{}
	dummy := &ListNode{Next: next}

	ost := 0
	for l1 != nil || l2 != nil || ost != 0 {
		sum := ost

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		node := &ListNode{Val: sum % 10}
		next.Next = node
		next = node

		if sum/10 != 0 {
			ost = sum / 10
		} else {
			ost = 0
		}
	}

	return dummy.Next.Next
}
