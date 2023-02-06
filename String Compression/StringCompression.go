package main

import "strconv"

func main() {

}

func stringCompression(chars []byte) int {
	left, right := 0, 0

	for right < len(chars) {
		count := 1
		for right+1 < len(chars) && chars[right] == chars[right+1] {
			count++
			right++
		}

		chars[left] = chars[right]
		left++

		if count > 1 {
			for _, v := range strconv.Itoa(count) {
				chars[left] = byte(v)
				left++
			}
		}
		right++
	}
	return left
}
