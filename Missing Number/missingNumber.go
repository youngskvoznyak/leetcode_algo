package main

func main() {

}

func missingNumber(nums []int) int {
	var allsum, arrsum int

	for i := 0; i < len(nums); i++ {
		arrsum += nums[i]
		allsum += i
	}
	return (allsum + len(nums)) - arrsum
}
