package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur := head

	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
