package main

func main() {

}

func moveZeros(nums []int) {
	pos := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos++
		}
	}
}
