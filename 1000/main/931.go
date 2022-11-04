package main

import "fmt"

//931. Minimum Falling Path Sum
//931. 下降路径最小和
var gMatrix [][]int
var gMin [][]int
var length int

func minFallingPathSum(matrix [][]int) int {
	gMatrix = matrix
	length = len(matrix)
	gMin = make([][]int, length-1)
	for i := 0; i < length-1; i++ {
		gMin[i] = make([]int, length)
		for j := 0; j < length; j++ {
			gMin[i][j] = 666666
		}
	}
	min := 666666
	for i := 0; i < length; i++ {
		if min > getMin(i, 0) {
			min = getMin(i, 0)
		}

	}
	return min
}

func getMin(x int, y int) int {
	if x < 0 || x >= length {
		return 666666
	}
	if y == length-1 {
		return gMatrix[y][x]
	}
	if gMin[y][x] != 666666 {
		return gMin[y][x]
	}
	min := getMin(x-1, y+1)
	if min > getMin(x, y+1) {
		min = getMin(x, y+1)
	}
	if min > getMin(x+1, y+1) {
		min = getMin(x+1, y+1)
	}
	gMin[y][x] = min + gMatrix[y][x]
	return gMin[y][x]
}

func main() {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}

	res := minFallingPathSum(matrix)

	fmt.Println(res)
}
