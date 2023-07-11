package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	var res [][]int
	// сотритруем исходный массив
	sort.Ints(nums)

	n := len(nums)

	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		// пропускаем дубликаты слева
		// num1Idx > 0 гарантирует, что эта проверка выполняется только начиная со 2-го элемента
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1
		num3Idx := n - 1

		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				res = append(res, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})
				num3Idx--

				// пропускаем дубликаты справа
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				num3Idx--
			} else {
				num2Idx++
			}
		}
	}
	return res
}
