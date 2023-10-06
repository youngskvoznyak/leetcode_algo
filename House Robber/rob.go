package main

func main() {

}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l1, l2 := nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		l2, l1 = max(l1+nums[i], l2), l2
	}
	return l2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
