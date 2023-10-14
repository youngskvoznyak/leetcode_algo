package main

func main() {

}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= len(s); i++ {
		s2 := s[i-2 : i]
		if "10" <= s2 && s2 <= "26" {
			dp[i] = dp[i-2]
		}

		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
	}
	return dp[len(s)]
}
