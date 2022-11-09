package main

import "fmt"

//886. Possible Bipartition
//886. 可能的二分法

var res bool
var graphInt [][]int
var visited, color []bool

func possibleBipartition(n int, dislikes [][]int) bool {
	if len(dislikes) == 0 {
		return true
	}
	graphInt = make([][]int, n+1)
	for i := 0; i < len(dislikes); i++ {
		graphInt[dislikes[i][0]] = append(graphInt[dislikes[i][0]], dislikes[i][1])
		graphInt[dislikes[i][1]] = append(graphInt[dislikes[i][1]], dislikes[i][0])
	}
	visited = make([]bool, n+1)
	color = make([]bool, n+1)
	res = true
	for i, _ := range graphInt {
		if visited[i] == false {
			traverse(i)
		}
	}

	return res
}

func traverse(index int) {
	if !res {
		return
	}
	visited[index] = true
	for _, i2 := range graphInt[index] {
		if visited[i2] == true {
			if color[index] == color[i2] {
				res = false
				return
			}
			continue
		} else {
			color[i2] = !color[index]
		}
		traverse(i2)
	}
}

func main() {
	//n := 4
	//dislikes := [][]int{{1, 2}, {1, 3}, {2, 4}}
	n := 10
	dislikes := [][]int{{1, 2}, {3, 4}, {5, 6}, {6, 7}, {8, 9}, {7, 8}}
	//n := 4
	//dislikes := [][]int{{1, 2}, {3, 4}, {1, 3}, {1, 4}}

	//n := 5
	//dislikes := [][]int{{1, 2}, {3, 4}, {4, 5}, {3, 5}}

	res := possibleBipartition(n, dislikes)

	fmt.Println(res)
}
