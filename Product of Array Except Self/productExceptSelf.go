package main

func main() {

}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums)) // создаем массив , где len равна длине исходного массива

	// проходимся по созданному массиву и присваеваем элементу по индексу единицу
	for i := range ans {
		ans[i] = 1
	}
	// выглядит так [1, 1, 1, 1, 1]

	// переменаня prefix нужна для префиксного умножения
	prefix := 1

	// проходимся по исходному массиву. элемент по индексу будет равен произведению prefix * nums[i]
	for i := range nums {
		ans[i] = prefix
		prefix *= nums[i]
	}
	// пример [1, 2 ,3 ,4]
	// ans =  [1, 1, 1, 1] => [1, 1, 2, 6]
	postfix := 1

	// идем с конца массива. теперь элемент массива равне ans[i] * postfix. а postfix будет равен postfix * nums[i]
	// ans = [1, 1, 2, 6] => [6] => [8,6] => [12, 8, 6] => [24, 12, 8, 6]
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= postfix
		postfix *= nums[i]
	}
	return ans
}
