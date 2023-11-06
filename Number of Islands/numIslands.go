package main

func main() {

}

func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	res := 0
	var visit func(grid [][]byte, r, c int)
	visit = func(grid [][]byte, r, c int) {
		if r < 0 || r >= row || c < 0 || c >= col || grid[r][c] == '0' {
			return
		}
		grid[r][c] = '0'
		visit(grid, r-1, c)
		visit(grid, r+1, c)
		visit(grid, r, c+1)
		visit(grid, r, c-1)
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' {
				visit(grid, r, c)
				res++
			}
		}
	}

	return res
}
