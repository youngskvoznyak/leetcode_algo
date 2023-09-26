package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1

	for l <= r {
		mid := l + (r-l)/2
		mid_val := matrix[mid/n][mid%n]

		if target == mid_val {
			return true
		} else if target > mid_val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
