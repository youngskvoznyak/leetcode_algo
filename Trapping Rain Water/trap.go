package main

func main() {

}

func trap(height []int) int {
	if height == nil {
		return 0
	}

	l, r, res := 0, len(height)-1, 0
	maxL, maxR := height[l], height[r]

	for l < r {
		if maxL < maxR {
			l++
			maxL = max(maxL, height[l])
			res += maxL - height[l]
		} else {
			r--
			maxR = max(maxR, height[r])
			res += maxR - height[r]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
