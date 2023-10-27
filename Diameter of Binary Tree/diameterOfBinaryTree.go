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

func getHeight(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	left := getHeight(root.Left, ans)
	right := getHeight(root.Right, ans)
	*ans = max(*ans, left+right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
