package main

func main() {

}

func topKFrequent(nums []int, k int) (res []int) {
	countMap := make(map[int]int)            // создаем мапу для подсчета количества элементов в массиве
	countSlice := make([][]int, len(nums)+1) // создаем массив бакетов для размещения в них (индекс равен количеству элемента) элементов

	// заполняем мапу элемент[количество]
	for _, num := range nums {
		if count, ok := countMap[num]; ok {
			countMap[num] = count + 1
		} else {
			countMap[num] = 1
		}
	}
	// заполняем массив бакетов. индекс = количеству. значение по индексу = элемент
	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	// проходимся с конца массива и добавляем в рез.массив самые частые элементы, пока их количество меньше k
	for i := len(countSlice) - 1; i > 0; i-- {
		res = append(res, countSlice[i]...)
		if len(res) == k {
			return
		}
	}
	return
}
