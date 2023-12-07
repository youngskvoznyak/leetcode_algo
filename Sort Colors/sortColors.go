package main

func main() {

}

func sortColors(nums []int) {
	l, r := 0, len(nums)-1

	for i := 0; i < r; {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			i++
			l++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			i++
		}
	}
}
