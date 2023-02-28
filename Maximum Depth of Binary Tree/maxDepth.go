package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxDepth(root *TreeNode) int {
	// if root == nil{
	//     return 0
	// }

	// if root.Left == nil && root.Right == nil{
	//     return 1
	// }

	// if root.Left == nil{
	//     return maxDepth(root.Right) + 1
	// }
	// if root.Right == nil{
	//     return maxDepth(root.Left) + 1
	// }
	// return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	lvl := 0
	for len(queue) != 0 {
		nxtlvl := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				nxtlvl = append(nxtlvl, node.Left)
			}
			if node.Right != nil {
				nxtlvl = append(nxtlvl, node.Right)
			}
		}
		queue = nxtlvl
		lvl++
	}
	return lvl
}

// func max(a, b int) int{
//     if a > b{
//         return a
//     }
//     return b
// }
