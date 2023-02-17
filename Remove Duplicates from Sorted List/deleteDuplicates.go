package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	for cur := head; cur != nil; cur = cur.Next {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
	}
	return head
}
