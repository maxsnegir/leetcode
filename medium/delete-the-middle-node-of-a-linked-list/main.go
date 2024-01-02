package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func deleteMiddle(head *ListNode) *ListNode {

	if head.Next == nil {
		return nil
	}

	var prevLeft *ListNode
	left := head
	right := head

	for right != nil && right.Next != nil {
		prevLeft = left
		left = left.Next
		right = right.Next.Next
	}

	if prevLeft != nil {
		prevLeft.Next = left.Next
	} else {
		left.Next = nil
	}

	return head
}
