package main

func main() {

}

func generateParentheses(n int) []string {
	var res []string
	helper(&res, n, 0, 0, "")
	return res
}

func helper(res *[]string, pairs, open, close int, curstr string) {
	if open == pairs && close == pairs {
		*res = append(*res, curstr)
		return
	}

	if open < pairs {
		helper(res, pairs, open+1, close, curstr+"(")
	}

	if close < open {
		helper(res, pairs, open, close+1, curstr+")")
	}
}
