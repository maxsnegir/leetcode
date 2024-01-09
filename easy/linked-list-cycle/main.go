package main

type ListNode struct {
	Next *ListNode
	Val  int
}

func main() {

}

func hasCycle(head *ListNode) bool {

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
