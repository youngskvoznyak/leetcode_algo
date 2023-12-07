package main

func main() {

}

func numSubarrayProductLessThanK(nums []int, k int) int {
	total := 0
	l := 0
	running := 1

	for r := 0; r < len(nums); r++ {
		running *= nums[r]

		for l < r && running >= k {
			running /= nums[l]
			l++
		}

		if running < k {
			total += r - l + 1
		}
	}
	return total
}
