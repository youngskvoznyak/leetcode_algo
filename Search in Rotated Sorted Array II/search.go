package main

func main() {

}

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return true
		}

		if nums[mid] == nums[l] && nums[r] == nums[mid] {
			l++
			r--
		} else if nums[l] <= nums[mid] {
			if target < nums[l] || target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target > nums[r] || target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}
