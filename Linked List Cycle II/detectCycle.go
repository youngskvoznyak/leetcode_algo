package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			slow = head
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
