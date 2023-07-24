package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxDepth(root *TreeNode) int {
	// 1. recursive solution
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

	// 2. iterative solution
	// if root == nil {
	// 	return 0
	// }
	// queue := []*TreeNode{root}
	// lvl := 0
	// for len(queue) != 0 {
	// 	nxtlvl := []*TreeNode{}
	// 	for _, node := range queue {
	// 		if node.Left != nil {
	// 			nxtlvl = append(nxtlvl, node.Left)
	// 		}
	// 		if node.Right != nil {
	// 			nxtlvl = append(nxtlvl, node.Right)
	// 		}
	// 	}
	// 	queue = nxtlvl
	// 	lvl++
	// }
	// return lvl
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
