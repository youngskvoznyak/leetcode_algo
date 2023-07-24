package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return areEqualTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func areEqualTree(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 == nil || n2 == nil {
		return false
	} else {
		return n1.Val == n2.Val && areEqualTree(n1.Left, n2.Left) && areEqualTree(n1.Right, n2.Right)
	}
}
