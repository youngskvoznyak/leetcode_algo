package main

func main() {

}

func maxArea(height []int) int {
	l, r, res := 0, len(height)-1, 0

	for l < r {
		// проверяем максимальную площадь чтобы не разлилась вода
		area := min(height[l], height[r]) * (r - l)

		if area > res {
			res = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
