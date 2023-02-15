package main

type NumArray struct {
	nums    []int
	preSums []int
}

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums))
	arr[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		arr[i] = nums[i] + arr[i-1]
	}

	return NumArray{nums: nums, preSums: arr}
}

func (this *NumArray) SumRange(left, right int) int {
	if left == 0 {
		return this.preSums[right]
	}
	return this.preSums[right] - this.preSums[left-1]
}
