package main

func main() {

}

func canJump(nums []int) bool {
	for pos, farthest := 0, 0; pos < len(nums); pos++ {
		if pos > farthest {
			return false
		}
		farthest = max(farthest, pos+nums[pos])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
