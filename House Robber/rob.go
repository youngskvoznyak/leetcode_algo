package main

func main() {

}

func rob(nums []int) int {
	prev_1, prev_2 := 0, 0

	for _, value := range nums {
		tmp := max(prev_1+value, prev_2)
		prev_1, prev_2 = prev_2, tmp
	}
	return prev_2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
