package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
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
