package main

func main() {

}

func countBits(n int) []int {
	memo := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		memo[i] = memo[i>>1] + (i % 2)
	}
	return memo[:n+1]
}
