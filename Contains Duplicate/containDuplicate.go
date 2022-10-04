package main

func main() {

}

func containsDuplicateOne(nums []int) bool {
	visit := make(map[int]bool)
	for _, v := range nums {
		if visit[v] {
			return true
		}
		visit[v] = true
	}
	return false
}

func containsDuplicatTwo(nums []int) bool {
	visit := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := visit[v]; ok {
			return true
		}
		visit[v] = struct{}{}
	}
	return false
}
