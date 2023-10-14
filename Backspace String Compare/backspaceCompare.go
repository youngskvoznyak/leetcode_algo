package main

func main() {

}

func backspaceCompare(s, t string) bool {
	s1 := getString([]byte(s))
	s2 := getString([]byte(t))

	return s1 == s2
}
func getString(s []byte) string {
	start, end := 0, 0

	for ; end < len(s); end++ {
		if s[end] == '#' {
			start--
			if start < 0 {
				start = 0
			}
		} else {
			s[start], s[end] = s[end], s[start]
			start++
		}
	}
	return string(s[:start])
}
