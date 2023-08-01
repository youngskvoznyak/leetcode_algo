package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxPathSum(root *TreeNode) int {
	res := []int{math.MinInt32}
	maxSum(root, res)
	return res[0]
}

func maxSum(root *TreeNode, arr []int) int {
	if root == nil {
		return 0
	}

	left := max(0, maxSum(root.Left, arr))
	right := max(0, maxSum(root.Right, arr))
	arr[0] = max(arr[0], root.Val+left+right)

	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
