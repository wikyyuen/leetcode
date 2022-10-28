package main

import "fmt"

//1020. Number of Enclaves
//飞地的数量
var m, n, total int
var visited [][]bool

//var dao [][]int

func numEnclaves(grid [][]int) int {
	n = len(grid)
	m = len(grid[0])
	if n == 1 || m == 1 {
		return 0
	}
	//dao = [][]int{}
	total = 0
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	//先填岛
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(grid, 0, i)
		dfs(grid, m-1, i)
	}
	//再数
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			//遍历内部
			if grid[j][i] == 1 && visited[j][i] == false {
				dfs2(grid, i, j)
			}
		}
	}
	return total
}

func dfs(grid [][]int, x int, y int) {
	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}

	if grid[y][x] == 0 {
		return
	}
	grid[y][x] = 0
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
}

func dfs2(grid [][]int, x int, y int) {
	if grid[y][x] == 0 {
		return
	}
	if visited[y][x] {
		return
	}
	visited[y][x] = true
	total++
	//dao = append(dao, []int{x, y})
	dfs2(grid, x+1, y)
	dfs2(grid, x-1, y)
	dfs2(grid, x, y+1)
	dfs2(grid, x, y-1)
}

func main() {
	//grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	//grid := [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	grid := [][]int{{0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {1, 1, 0, 0, 0, 1, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {0, 1, 1, 0, 0, 0, 1, 0, 1, 0}, {0, 1, 1, 1, 1, 1, 0, 0, 1, 0}, {0, 0, 1, 0, 1, 1, 1, 1, 0, 1}, {0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, {0, 0, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 1, 0, 0, 0, 1}}
	//grid := [][]int{{0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1}, {1, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0}, {1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0}, {1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0}, {0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0}, {1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0}, {1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1}}
	result := numEnclaves(grid)

	fmt.Println(result)
	//fmt.Println(dao)
}
