package main

func main() {

}

func maxProfit(prices []int) int {
	current_min, profit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		profit = max(profit, prices[i]-current_min)
		current_min = min(current_min, prices[i])
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
