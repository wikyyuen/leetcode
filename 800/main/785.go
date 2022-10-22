package main

import "fmt"

//785. Is Graph Bipartite?
//判断二分图
var res bool
var graphInt [][]int
var visited, color []bool

func isBipartite(graph [][]int) bool {
	graphInt = graph
	visited = make([]bool, len(graph))
	color = make([]bool, len(graph))
	res = true
	for i, _ := range graphInt {
		if visited[i] == false {
			traverse2(i)
		}
	}

	return res
}

func traverse2(index int) {
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
		traverse2(i2)
	}
}

func main() {
	//graph := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	//graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	//graph := [][]int{{4}, {}, {4}, {4}, {0, 2, 3}}
	graph := [][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}

	result := isBipartite(graph)

	fmt.Println(result)
}
