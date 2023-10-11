package main

func main() {

}

func longestPalindrome(s string) string {
	res_len := 0
	var res string

	for i := 0; i < len(s); i++ {
		// odd length
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > res_len {
				res_len = r - l + 1
				res = s[l : r+1]
			}
			l--
			r++
		}

		// even length
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > res_len {
				res_len = r - l + 1
				res = s[l : r+1]
			}
			l--
			r++
		}
	}
	return res
}
