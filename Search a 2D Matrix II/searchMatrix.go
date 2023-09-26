package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	r, c := 0, len(matrix[0])-1

	for r < len(matrix) && c >= 0 {
		if target == matrix[r][c] {
			return true
		} else if target < matrix[r][c] {
			c--
		} else {
			r++
		}
	}
	return false
}
