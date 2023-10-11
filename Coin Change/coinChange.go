package main

import "math"

func main() {

}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = math.MaxInt16
	}

	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] < math.MaxInt16 {
		return dp[amount]
	} else {
		return -1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
