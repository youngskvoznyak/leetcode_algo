package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	curPath := []int{}

	var dfs func(*TreeNode, []int, int)

	dfs = func(root *TreeNode, curPath []int, curSum int) {
		if root == nil {
			return
		}
		curPath = append(curPath, root.Val)

		curSum += root.Val

		if root.Left == nil && root.Right == nil {
			if curSum == targetSum {
				curSol := append([]int{}, curPath...)
				ans = append(ans, curSol)
				curPath = curPath[:len(curPath)-1]
				return
			}
		}
		dfs(root.Left, curPath, curSum)
		dfs(root.Right, curPath, curSum)

		curPath = curPath[:len(curPath)-1]
		return
	}

	dfs(root, curPath, 0)

	return ans
}
