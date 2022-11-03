package main

import "fmt"

//174. Dungeon Game
//174. 地下城游戏

var gDungeon, dpList [][]int
var h, w int

func calculateMinimumHP(dungeon [][]int) int {
	gDungeon = dungeon
	h = len(dungeon)
	w = len(dungeon[0])
	dpList = make([][]int, h)
	for i := 0; i < h; i++ {
		dpList[i] = make([]int, w)
		for j := 0; j < w; j++ {
			dpList[i][j] = -999999
		}
	}
	res := dp(0, 0)
	if res < 0 {
		res = 0 - res
	} else {
		res = 0
	}
	return res + 1
}

func dp(x int, y int) int {
	if dpList[y][x] != -999999 {
		return dpList[y][x]
	}
	max := 0
	if x == w-1 || y == h-1 {
		if x == w-1 && y == h-1 {
			max = 0
		} else if y == h-1 {
			max = dp(x+1, y)
		} else {
			max = dp(x, y+1)
		}
	} else {
		max = dp(x+1, y)
		down := dp(x, y+1)
		if max < down {
			max = down
		}
	}
	dpList[y][x] = gDungeon[y][x] + max
	if dpList[y][x] > 0 {
		dpList[y][x] = 0
	}
	return dpList[y][x]
}

func main() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}

	res := calculateMinimumHP(dungeon)

	fmt.Println(res)
}
