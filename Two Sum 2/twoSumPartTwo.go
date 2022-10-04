package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	sum := numbers[l] + numbers[r]
	for sum != target {
		if sum < target {
			l++
		} else {
			r--
		}
		sum = numbers[l] + numbers[r]
	}
	return []int{l + 1, r + 1}
}
