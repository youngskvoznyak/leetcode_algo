package main

import "math"

func main() {

}

func maxProfit(prices []int) int {
	// initialization
	cool_down, sell, hold := 0, 0, math.MinInt64

	for _, stock_price_of_Day_i := range prices {

		prev_cool_down, prev_sell, prev_hold := cool_down, sell, hold

		// Max profit of cooldown on Day i comes from either cool down of Day_i-1, or sell out of Day_i-1 and today Day_i is cooling day
		cool_down = max(prev_cool_down, prev_sell)

		// Max profit of sell on Day_i comes from hold of Day_i-1 and sell on Day_i
		sell = prev_hold + stock_price_of_Day_i

		// Max profit of hold on Day_i comes from either hold of Day_i-1, or cool down on Day_i-1 and buy on Day_i
		hold = max(prev_hold, prev_cool_down-stock_price_of_Day_i)

	}

	// The state of max profit must be either sell or cool down
	return max(sell, cool_down)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
