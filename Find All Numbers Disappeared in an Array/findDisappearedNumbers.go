package main

func main() {

}

func findDisappearedNumbers(nums []int) []int {
	i := 0

	for i < len(nums) {
		pos := nums[i] - 1

		if nums[i] != nums[pos] {
			nums[pos], nums[i] = nums[i], nums[pos]
		} else {
			i++
		}
	}

	var res []int
	for i := range nums {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
