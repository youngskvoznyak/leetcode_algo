package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return areSubTrees(root, subRoot) || areSubTrees(root.Left, subRoot) || areSubTrees(root.Right, subRoot)

}

func areSubTrees(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 == nil || n2 == nil {
		return false
	} else {
		return n1.Val == n2.Val && areSubTrees(n1.Left, n2.Left) && areSubTrees(n2.Right, n2.Right)
	}
}
