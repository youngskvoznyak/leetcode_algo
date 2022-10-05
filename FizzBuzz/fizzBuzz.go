package main

import "strconv"

func main() {}

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		temp := ""
		flag := 0
		if i%3 == 0 {
			temp += "Fizz"
			flag = 1
		}
		if i%5 == 0 {
			temp += "Buzz"
			flag = 1
		}
		if flag == 0 {
			temp += strconv.Itoa(i)
		}
		res = append(res, temp)
	}
	return res
}
