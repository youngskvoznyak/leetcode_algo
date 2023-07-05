package main

func main() {

}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	longest := 0

	for num := range set {
		if set[num-1] {
			continue
		}

		curlen := 0

		for set[num] {
			curlen++
			num++
		}
		longest = max(longest, curlen)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
