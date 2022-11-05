package main

import "fmt"

//130. Surrounded Regions
//被围绕的区域

var visited [][]bool
var m, n int
var gBoard [][]byte

func solve(board [][]byte) {
	n = len(board)
	m = len(board[0])
	gBoard = board

	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	//先找到所以没有被围绕的
	for i := 0; i < m; i++ {
		dp(i, 0)
		dp(i, n-1)
	}
	for i := 1; i < n-1; i++ {
		dp(0, i)
		dp(m-1, i)
	}

	//开始填充
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if board[i][j] == 79 && visited[i][j] == false {
				board[i][j] = 88
			}
		}
	}
}

func dp(x int, y int) {
	if visited[y][x] == true || gBoard[y][x] == 88 {
		return
	}
	visited[y][x] = true
	if x > 0 {
		dp(x-1, y)
	}
	if x < m-1 {
		dp(x+1, y)
	}
	if y > 0 {
		dp(x, y-1)
	}
	if y < n-1 {
		dp(x, y+1)
	}
}

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}

	solve(board)

	fmt.Println(board)
}
