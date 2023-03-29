package main

func main() {

}

func setZeroes(matrix [][]int) {
	sizeI, sizeJ := len(matrix), len(matrix[0])
	markI, markJ := -1, -1

	for i := 0; i < sizeI; i++ {
		for j := 0; j < sizeJ; j++ {
			if matrix[i][j] == 0 {
				if markI == -1 {
					markI = i
					markJ = j
				} else {
					matrix[i][markJ] = 0
					matrix[markI][j] = 0
				}
			}
		}
	}
	if markI == -1 {
		return
	}

	for i := 0; i < sizeI; i++ {
		for j := 0; j < sizeJ; j++ {
			if i != markI && j != markJ {
				if matrix[markI][j] == 0 || matrix[i][markJ] == 0 {
					matrix[i][j] = 0
				}
			}
		}
	}
	for i := 0; i < sizeI; i++ {
		matrix[i][markJ] = 0
	}
	for j := 0; j < sizeJ; j++ {
		matrix[markI][j] = 0
	}
}
