package main

func main() {

}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range nums {
		remaining := target - v
		if ind, ok := seen[remaining]; ok {
			return []int{ind, i}
		}
		seen[v] = i
	}
	return []int{}
}
