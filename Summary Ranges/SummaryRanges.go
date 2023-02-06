package main

import "fmt"

func main() {

}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var ans []string
	start, end := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			end = nums[i]
		} else {
			if start != end {
				ans = append(ans, fmt.Sprintf("%d->%d", start, end))
			} else {
				ans = append(ans, fmt.Sprint(start))
			}
			start, end = nums[i], nums[i]
		}
	}

	if start != end {
		ans = append(ans, fmt.Sprintf("%d->%d", start, end))
	} else {
		ans = append(ans, fmt.Sprint(start))
	}

	return ans
}
