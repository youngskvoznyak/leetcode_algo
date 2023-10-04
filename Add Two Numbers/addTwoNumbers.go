package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	dummy, carry := new(ListNode), 0

	cur := dummy

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: carry % 10}
		carry /= 10
		cur = cur.Next
	}
	return dummy.Next
}
