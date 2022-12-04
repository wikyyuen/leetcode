package main

import "fmt"

//1905. Count Sub Islands
//1905. 统计子岛屿

var visited [][]bool
var gGrid1, gGrid2 [][]int
var m, n int
var isInclude bool

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	n = len(grid2)
	m = len(grid2[0])
	gGrid2 = grid2
	gGrid1 = grid1
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			isInclude = true
			if visited[i][j] || gGrid2[i][j] == 0 {
				continue
			}
			dp2(j, i)
			if isInclude {
				total++
			}
		}
	}
	return total
}

func dp2(x int, y int) {
	if visited[y][x] || gGrid2[y][x] == 0 {
		return
	}
	visited[y][x] = true
	//上下左右
	if y > 0 {
		dp2(x, y-1)
	}
	if y < n-1 {
		dp2(x, y+1)
	}
	if x > 0 {
		dp2(x-1, y)
	}
	if x < m-1 {
		dp2(x+1, y)
	}
	if isInclude && gGrid1[y][x] == 0 {
		isInclude = false
	}
}

func main() {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 1}}

	res := countSubIslands(grid1, grid2)

	fmt.Println(res)
}
