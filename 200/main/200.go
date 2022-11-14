package main

import "fmt"

//200. Number of Islands
//200. 岛屿数量

var visited [][]bool
var m, n, sum int
var gGrid [][]byte

func numIslands(grid [][]byte) int {
	gGrid = grid
	n = len(grid)
	m = len(grid[0])
	sum = 0
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp2(j, i, true)
		}
	}
	return sum
}

func dp2(x int, y int, isStart bool) {
	if visited[y][x] || gGrid[y][x] == '0' {
		return
	}
	if isStart {
		sum++
	}
	visited[y][x] = true
	//上下左右
	if y > 0 {
		dp2(x, y-1, false)
	}
	if y < n-1 {
		dp2(x, y+1, false)
	}
	if x > 0 {
		dp2(x-1, y, false)
	}
	if x < m-1 {
		dp2(x+1, y, false)
	}
}

func main() {
	grid := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}

	res := numIslands(grid)

	fmt.Println(res)
}
