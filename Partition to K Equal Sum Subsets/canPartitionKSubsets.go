package main

func main() {

}

var memo map[int]bool

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}

	memo = make(map[int]bool)

	used := 0
	target := sum / k
	return backtrack(k, 0, nums, 0, used, target)
}

func backtrack(k int, bucketSum int, nums []int, startIdx int, used int, target int) bool {
	// base case, if all bucket has been filled, return true
	if k == 0 {
		return true
	}

	// if current bucket is full
	if bucketSum == target {
		// go to the next bucket, and record the used status
		result := backtrack(k-1, 0, nums, 0, used, target)
		memo[used] = result
		return result
	}

	// means this status has been calcuated before
	if res, ok := memo[used]; ok {
		return res
	}

	for i := startIdx; i < len(nums); i++ {
		// if this number has been put into another bucket
		if ((used >> i) & 1) == 1 {
			continue
		}

		// if cannot hold this number, continue
		if nums[i]+bucketSum > target {
			continue
		}

		// put this number into this bucket
		used |= 1 << i
		bucketSum += nums[i]

		if backtrack(k, bucketSum, nums, i+1, used, target) {
			return true
		}

		// unset tthe i-th bit
		used ^= 1 << i
		bucketSum -= nums[i]
	}

	// all the number cannot satisfy
	return false
}
