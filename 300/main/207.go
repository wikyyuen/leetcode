package main

import "fmt"

//207. Course Schedule
//207. 课程表

var res bool
var graphInt [][]int
var visited, onPath []bool

func canFinish(numCourses int, prerequisites [][]int) bool {
	graphInt = make([][]int, numCourses)
	for _, i2 := range prerequisites {
		graphInt[i2[1]] = append(graphInt[i2[1]], i2[0])
	}
	visited = make([]bool, numCourses)
	onPath = make([]bool, numCourses)
	res = true
	for i, _ := range graphInt {
		if visited[i] == false {
			traverse2(i)
		}
	}

	return res
}

func traverse2(index int) {
	if onPath[index] {
		res = false
	}

	if visited[index] || onPath[index] {
		return
	}
	visited[index] = true
	onPath[index] = true
	for _, i2 := range graphInt[index] {
		traverse2(i2)
	}
	onPath[index] = false
}

func main() {
	numCoures := 2
	//prerequisites := [][]int{{1, 0}}
	prerequisites := [][]int{{1, 0}, {0, 1}}
	result := canFinish(numCoures, prerequisites)

	fmt.Println(result)
}
