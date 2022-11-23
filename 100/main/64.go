package main

import "fmt"

//64. Minimum Path Sum
//64. 最小路径和

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dpList := make([][]int, n)
	for i := 0; i < n; i++ {
		dpList[i] = make([]int, m)
	}
	dpList[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dpList[0][i] = dpList[0][i-1] + grid[0][i]
	}
	for i := 1; i < n; i++ {
		dpList[i][0] = dpList[i-1][0] + grid[i][0]
	}
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dpList[i][j] = min(dpList[i-1][j], dpList[i][j-1]) + grid[i][j]
		}
	}
	return dpList[n-1][m-1]
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	res := minPathSum(grid)

	fmt.Println(res)
}
