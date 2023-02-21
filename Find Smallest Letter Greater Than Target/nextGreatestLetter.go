package main

func main() {

}

func nextGreatestLetter(letters []byte, target byte) byte {
	res := byte(letters[0])
	l, r := 0, len(letters)-1

	for l <= r {
		mid := (l + r) / 2
		if letters[mid] < target {
			l = mid + 1
		} else if letters[mid] == target {
			l = mid + 1
		} else {
			res = letters[0]
			r = mid - 1
		}
	}
	return res
}
