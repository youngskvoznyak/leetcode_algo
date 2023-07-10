package main

func main() {

}

func longestConsecutive(nums []int) int {
	// cоздаем набор из массива nums.
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	// переменная для ответа
	longest := 0

	// проходимся по набору
	for n := range set {
		// мы проверяем, есть ли n-1 в наборе.
		// если есть, то n не является началом последовательности, и мы переходим к следующему числу.
		if _, ok := set[n-1]; !ok {
			// в противном случае мы увеличиваем n в цикле,
			// чтобы увидеть, сохранено ли следующее последовательное значение в числах.
			seqLen := 1
			for {
				if _, ok := set[n+seqLen]; ok {
					seqLen++
					continue
				}
				// обновляем переменную
				longest = max(longest, seqLen)
				break
			}
		}
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
