package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reorderList(head *ListNode) {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversed := reverse(slow.Next)
	slow.Next = nil

	curr := head

	for curr != nil && reversed != nil {
		next := curr.Next
		revNext := reversed.Next
		curr.Next = reversed
		reversed.Next = next
		curr = next
		reversed = revNext
	}
}

func reverse(node *ListNode) *ListNode {
	var prev, curr *ListNode = nil, node

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
