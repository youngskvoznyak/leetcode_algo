package main

func main() {

}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)

		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		row[0], row[i] = 1, 1
		res[i] = row
	}

	return res
}
