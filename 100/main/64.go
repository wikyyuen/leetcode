package main

import "fmt"

//64. Minimum Path Sum
//64. 最小路径和

var dpList [][]int
var m, n int

func minPathSum(grid [][]int) int {
	n = len(grid)
	m = len(grid[0])
	dpList = make([][]int, n)
	for i := 0; i < n; i++ {
		dpList[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dpList[i][j] = -1
		}
	}
	dpList[n-1][m-1] = grid[n-1][m-1]
	return dp5(grid, 0, 0)
}

func dp5(grid [][]int, x int, y int) int {
	if dpList[y][x] != -1 {
		return dpList[y][x]
	}
	min := 999999
	if x < m-1 {
		temp1 := dp5(grid, x+1, y)
		if min > temp1 {
			min = temp1
		}
	}
	if y < n-1 {
		temp2 := dp5(grid, x, y+1)
		if min > temp2 {
			min = temp2
		}
	}

	dpList[y][x] = grid[y][x] + min
	return dpList[y][x]
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	res := minPathSum(grid)

	fmt.Println(res)
}
