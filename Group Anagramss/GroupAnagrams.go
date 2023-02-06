package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, s := range strs {
		var freq [26]int

		for _, c := range s {
			freq[c-'a']++
		}
		m[freq] = append(m[freq], s)
	}

	var groups [][]string

	for _, v := range m {
		groups = append(groups, v)
	}
	return groups
}
