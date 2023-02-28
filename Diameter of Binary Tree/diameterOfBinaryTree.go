package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	getHeight(root, &ans)
	return ans
}

func getHeight(node *TreeNode, ans *int) int {
	if node == nil {
		return -1
	}

	leftH := getHeight(node.Left, ans) + 1
	rightH := getHeight(node.Right, ans) + 1
	path := leftH + rightH
	*ans = max(*ans, path)
	return max(leftH, rightH)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
