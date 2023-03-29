package main

func main() {

}

func findDuplicates(nums []int) []int {
	var ans []int

	for _, v := range nums {
		m := abs(v)
		if nums[m-1] < 0 {
			ans = append(ans, m)
		} else {
			nums[m-1] *= -1
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
