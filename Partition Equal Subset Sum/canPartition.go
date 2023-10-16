package main

func main() {

}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([]bool, sum, sum) // array to mark reachable numbers
	dp[0] = true
	for _, n := range nums {
		if n <= sum { // to skip too big numbers
			if dp[sum-n] == true { // we found our sum
				return true
			}
			for j := sum - n - 1; j >= 0; j-- { // we loop in opposite direction, because we don't want to check index and then loop over it
				if dp[j] == true {
					dp[j+n] = true
				}
			}
		}
	}
	return false
}
