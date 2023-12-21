package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func lowesCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p && root == q {
		return root
	}

	l := lowesCommonAncestor(root.Left, p, q)
	r := lowesCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}
	return r
}
