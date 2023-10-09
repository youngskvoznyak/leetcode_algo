package main

func main() {

}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	return max(helper(nums, 0, len(nums)-2), helper(nums, 1, len(nums)-1))
}

func helper(nums []int, start, end int) int {
	rob1, rob2 := 0, 0

	for i := start; i <= end; i++ {
		tmp := max(rob1+nums[i], rob2)
		rob1, rob2 = rob2, tmp
	}
	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
