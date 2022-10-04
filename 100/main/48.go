package main

import "fmt"

//48. Rotate Image

func rotate(matrix [][]int) [][]int {
	length := len(matrix)
	half := length / 2
	temp := make([]int, length)
	//上下调转
	for i := 0; i < half; i++ {
		temp = matrix[i]
		matrix[i] = matrix[length-i-1]
		matrix[length-i-1] = temp
	}
	//沿斜线翻转
	tempInt := 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			tempInt = matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tempInt
		}
	}
	return matrix
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	result := rotate(matrix)

	fmt.Println(result)
}
