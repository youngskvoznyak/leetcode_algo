package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head

	for cur != nil {
		var tmp *ListNode = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}
