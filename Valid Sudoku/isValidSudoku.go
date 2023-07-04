package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	hashmap := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row := i
			col := j

			cur_val := string(board[i][j])

			if cur_val == "." {
				continue
			}
			_, ok1 := hashmap[cur_val+"found in row"+string(row)]
			_, ok2 := hashmap[cur_val+"found in column"+string(col)]
			_, ok3 := hashmap[cur_val+"found in grid"+string(i/3)+"-"+string(j/3)]

			if ok1 || ok2 || ok3 {
				return false
			} else {
				hashmap[cur_val+"found in row"+string(row)] = true
				hashmap[cur_val+"found in column"+string(col)] = true
				hashmap[cur_val+"found in grid"+string(i/3)+"-"+string(j/3)] = true
			}
		}
	}
	return true
}
