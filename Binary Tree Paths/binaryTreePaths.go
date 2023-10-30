package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	path(root, "", &ans)
	return ans
}

func path(root *TreeNode, prefix string, ans *[]string) {
	if root == nil {
		return
	}

	if len(prefix) == 0 {
		prefix += strconv.Itoa(root.Val)
	} else {
		prefix += "->" + strconv.Itoa(root.Val)
	}

	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, prefix)
		return
	}
	path(root.Left, prefix, ans)
	path(root.Right, prefix, ans)
}
