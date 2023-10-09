package main

func main() {

}

func findTargetSumWays(nums []int, target int) int {
	dp := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		dp2 := make(map[int]int)
		for currSum, currCount := range dp {
			for sign := -1; sign <= 1; sign += 2 {
				dp2[currSum+sign*nums[i]] += currCount
			}
		}
		dp = dp2
	}
	return dp[target]
}
