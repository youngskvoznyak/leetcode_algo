package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	queue := []*TreeNode{root}
	cnt := 0
	n := len(queue)

	for n != 0 {
		cnt++

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				if node.Left == nil && node.Right == nil {
					return cnt
				}
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		n = len(queue)
	}
	return 0
}
