package main

func main() {

}

func longestSubarray(nums []int) int {
	l, r, k := 0, 0, 1

	for r < len(nums) {
		if nums[r] == 0 {
			k--
		}

		if k < 0 {
			if nums[l] == 0 {
				k++
			}
			l++
		}
		r++
	}
	return r - l - 1
}
