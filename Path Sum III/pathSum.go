package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func pathSum(root *TreeNode, targetSum int) int {
	return dfs(root, targetSum, []int{})
}

func dfs(node *TreeNode, targetSum int, path []int) int {
	if node == nil {
		return 0
	}

	path = append(path, node.Val)
	pathCount, pathSum := 0, 0

	for i := len(path) - 1; i >= 0; i-- {
		pathSum += path[i]
		if pathSum == targetSum {
			pathCount++
		}
	}
	pathCount += dfs(node.Left, targetSum, path)
	pathCount += dfs(node.Right, targetSum, path)

	path = path[:len(path)-1]

	return pathCount
}
