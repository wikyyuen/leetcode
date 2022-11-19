package main

import "fmt"

//695. Max Area of Island
//695. 岛屿的最大面积

var visited [][]bool
var gGrid [][]int
var m, n int

func maxAreaOfIsland(grid [][]int) int {
	n = len(grid)
	m = len(grid[0])
	gGrid = grid
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			total := 0
			dp2(j, i, &total)
			if total > max {
				max = total
			}
		}
	}
	return max
}

func dp2(x int, y int, total *int) {
	if visited[y][x] || gGrid[y][x] == 0 {
		return
	}
	*total++
	visited[y][x] = true
	//上下左右
	if y > 0 {
		dp2(x, y-1, total)
	}
	if y < n-1 {
		dp2(x, y+1, total)
	}
	if x > 0 {
		dp2(x-1, y, total)
	}
	if x < m-1 {
		dp2(x+1, y, total)
	}
}

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

	res := maxAreaOfIsland(grid)

	fmt.Println(res)
}
