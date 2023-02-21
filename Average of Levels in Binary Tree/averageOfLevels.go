package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	ans := []float64{}

	if root == nil {
		return ans
	}

	levelsOfNodes := []*TreeNode{root}

	for len(levelsOfNodes) != 0 {
		sum := 0
		count := 0
		nextLevel := []*TreeNode{}

		for _, node := range levelsOfNodes {
			sum += node.Val
			count += 1

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		levelsOfNodes = nextLevel
		ans = append(ans, float64(sum)/float64(count))
	}

	return ans
}
