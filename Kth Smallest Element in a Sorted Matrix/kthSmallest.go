package main

func main() {

}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	min, max := matrix[0][0], matrix[n-1][n-1]

	for min != max {
		mid := min + (max-min)/2
		count := countLessOrEqual(matrix, mid)
		if count < k {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}

func countLessOrEqual(matrix [][]int, mid int) int {
	count := 0
	row, col := 0, len(matrix)-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] <= mid {
			count += col + 1
			row++
		} else {
			col--
		}
	}
	return count
}
