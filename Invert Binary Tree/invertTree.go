package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// invertTree(node.Left)
	// invertTree(node.Right)
	// node.Left, node.Right = node.Right, node.Left
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return root
}
