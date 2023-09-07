package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{0, head}

	leftprev, current := dummy, head

	for i := 0; i < left-1; i++ {
		leftprev, current = current, current.Next
	}

	var prev *ListNode = nil

	for i := 0; i < (right-left)+1; i++ {
		tmp := current.Next
		current.Next = prev
		prev, current = current, tmp
	}

	leftprev.Next.Next = current
	leftprev.Next = prev

	return dummy.Next
}
