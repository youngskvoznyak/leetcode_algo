package main

func main() {

}

func countSubstrings(s string) int {
	res := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[r] == s[l] {
			res++
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[r] == s[l] {
			res++
			l--
			r++
		}
	}
	return res
}
