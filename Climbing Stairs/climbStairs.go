package main

func main() {

}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
