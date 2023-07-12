package main

import "strconv"

func main() {

}

func evalRPN(tokens []string) int {
	var stack []int
	var a, b int

	for _, v := range tokens {
		switch v {
		case "+":
			a, b, stack = getAndPopTwoLast(stack)
			stack = append(stack, (a + b))

		case "-":
			a, b, stack = getAndPopTwoLast(stack)
			stack = append(stack, (a - b))
		case "*":
			a, b, stack = getAndPopTwoLast(stack)
			stack = append(stack, (a * b))
		case "/":
			a, b, stack = getAndPopTwoLast(stack)
			stack = append(stack, (a / b))
		default:
			i, _ := strconv.Atoi(v)
			stack = append(stack, i)
		}
	}
	return stack[0]
}

func getAndPopTwoLast(stack []int) (int, int, []int) {
	a := stack[len(stack)-2]
	b := stack[len(stack)-1]
	stack = stack[:len(stack)-2]

	return a, b, stack
}
