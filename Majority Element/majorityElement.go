package main

func main() {

}

func majorityElement(nums []int) int {
	cand := 0
	cnt := 0

	for _, v := range nums {
		if cand == v {
			cnt++
		} else {
			if cnt == 0 {
				cand = v
				cnt = 1
			} else {
				cnt--
			}
		}
	}
	return cand
}
