package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func splitListToParts(head *ListNode, k int) []*ListNode {
	len := getLength(head)
	baseSize, extra := len/k, len%k
	var res []*ListNode

	cur := head
	for i := 0; i < k; i++ {
		partSize := baseSize
		if extra > 0 {
			partSize++
			extra--
		}

		var partHead, partTail *ListNode
		for j := 0; j < partSize; j++ {
			if partHead == nil {
				partHead, partTail = cur, cur
			} else {
				partTail.Next = cur
				partTail = partTail.Next
			}

			if cur != nil {
				cur = cur.Next
			}
		}
		if partTail != nil {
			partTail.Next = nil
		}
		res = append(res, partHead)
	}
	return res
}

func getLength(head *ListNode) int {
	current, length := head, 0

	for current != nil {
		length++
		current = current.Next
	}
	return length

}
