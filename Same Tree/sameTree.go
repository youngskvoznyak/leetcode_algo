package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sameTree(p, q *TreeNode) bool {
	// if p == nil && q == nil {
	// 	return true
	// }

	// if p == nil || q == nil || p.Val != q.Val {
	// 	return false
	// }

	// return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Left)

	pQueue := []*TreeNode{p}
	qQueue := []*TreeNode{q}

	for len(pQueue) != 0 && len(qQueue) != 0 {
		pNode := pQueue[0]
		pQueue = pQueue[1:]

		qNode := qQueue[0]
		qQueue = qQueue[1:]

		if pNode == nil && qNode == nil {
			continue
		}
		if pNode != nil && qNode == nil || pNode == nil && qNode != nil {
			return false
		}
		if pNode.Val != qNode.Val {
			return false
		}
		pQueue = append(pQueue, pNode.Left, pNode.Right)
		qQueue = append(qQueue, qNode.Left, qNode.Right)
	}
	if len(qQueue) == 0 && len(pQueue) == 0 {
		return true
	}
	return false
}
