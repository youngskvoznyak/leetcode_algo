package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	prev, cur := head, head

	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}

		cur = cur.Next
	}
	if head.Val == val {
		head = head.Next
	}

	return head
}
