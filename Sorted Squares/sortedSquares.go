package main

func main() {

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	ans := make([]int, n)
	for i := r; i >= 0; i-- {
		square := 0
		if abs(nums[l]) < abs(nums[r]) {
			square = nums[r]
			r--
		} else {
			square = nums[l]
			l++
		}
		ans[i] = square * square
	}
	return ans
}
