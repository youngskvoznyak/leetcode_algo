package main

func main() {

}

func convert(original []int, m, n int) [][]int {
	var ans [][]int

	if len(original) == n*m {
		for i := 0; i < m; i++ {
			ans = append(ans, original[i*n:(i+1)*n])
		}
	}
	return ans
}
