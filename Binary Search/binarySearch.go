package main

func main() {}

func binarySearch(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		mid := (l + h) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
