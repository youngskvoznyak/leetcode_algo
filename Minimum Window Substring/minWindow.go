package main

func main() {

}

func minWindow(s string, t string) string {
	start, end := 0, 0
	targetCharFreq := make(map[uint8]int)
	currentCharFreq := make(map[uint8]int)
	distinctCharCount := 0
	minSubstring := ""

	for index := range t {
		targetCharFreq[t[index]]++
	}

	for end < len(s) {
		currentCharFreq[s[end]]++
		if targetCharFreq[s[end]] != 0 &&
			targetCharFreq[s[end]] == currentCharFreq[s[end]] {
			distinctCharCount++
		}
		for distinctCharCount == len(targetCharFreq) {
			if minSubstring == "" {
				minSubstring = s[start : end+1]
			}
			if end-start+1 < len(minSubstring) {
				minSubstring = s[start : end+1]
			}
			currentCharFreq[s[start]]--
			if currentCharFreq[s[start]] < targetCharFreq[s[start]] {
				distinctCharCount--
			}
			start++
		}
		end++
	}
	return minSubstring
}
