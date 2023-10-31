package main

func main() {

}

var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func pacificAtlantic(heights [][]int) [][]int {
	pacificVisited := make([][]bool, len(heights))
	atlanticVisited := make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		pacificVisited[i] = make([]bool, len(heights[0]))
		atlanticVisited[i] = make([]bool, len(heights[0]))
	}

	// pacific - left edge
	for i := 0; i < len(heights); i++ {
		dfs(heights, &pacificVisited, i, 0)
	}
	// pacific - top edge
	for i := 1; i < len(heights[0]); i++ {
		dfs(heights, &pacificVisited, 0, i)
	}

	// atlantic - right edge
	for i := 0; i < len(heights); i++ {
		dfs(heights, &atlanticVisited, i, len(heights[0])-1)
	}
	// atlantic - bottom edge
	for i := 0; i < len(heights[0])-1; i++ {
		dfs(heights, &atlanticVisited, len(heights)-1, i)
	}

	ans := make([][]int, 0)

	// find cells where both visited
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if pacificVisited[i][j] && atlanticVisited[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func dfs(grid [][]int, visited *[][]bool, i, j int) {
	if (*visited)[i][j] {
		return
	}

	(*visited)[i][j] = true

	for _, dir := range dirs {
		ny, nx := i+dir[0], j+dir[1]
		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[0]) {
			if grid[ny][nx] >= grid[i][j] {
				dfs(grid, visited, ny, nx)
			}
		}
	}
}
