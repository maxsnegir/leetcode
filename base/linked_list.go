package base

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
