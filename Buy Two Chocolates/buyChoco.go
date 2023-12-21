package main

func main() {

}

func buyChoco(prices []int, money int) int {
	min1, min2 := prices[0], prices[1]
	if min2 < min1 {
		min1, min2 = min2, min1
	}

	for i := 2; i < len(prices); i++ {
		if prices[i] < min1 {
			min2 = min1
			min1 = prices[i]
		} else if prices[i] < min2 {
			min2 = prices[i]
		}
	}

	p := min1 + min2
	if p <= money {
		return money - p
	}
	return money
}
