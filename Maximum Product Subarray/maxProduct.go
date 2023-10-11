package main

func main() {

}

func maxProduct(nums []int) int {
	maxVal, minVal, maxProd := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxVal, minVal = minVal, maxVal
		}
		minVal = min(nums[i], minVal*nums[i])
		maxVal = max(nums[i], maxVal*nums[i])

		if maxVal > maxProd {
			maxProd = maxVal
		}
	}
	return maxProd
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
