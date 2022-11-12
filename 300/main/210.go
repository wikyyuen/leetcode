package main

import "fmt"

//210. Course Schedule II
//210. 课程表 II

var visited, onPath []bool
var hasCycle bool
var list []int

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph(numCourses, prerequisites)
	visited = make([]bool, numCourses)
	onPath = make([]bool, numCourses)
	list = []int{}
	hasCycle = false

	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}
	if hasCycle {
		return []int{}
	}
	reverse(&list)
	return list

}

func reverse(nums *[]int) {
	length := len(*nums)
	for i := 0; i < length/2; i++ {
		(*nums)[i], (*nums)[length-1-i] = (*nums)[length-1-i], (*nums)[i]
	}
}

func traverse(graph [][]int, index int) {
	if onPath[index] {
		hasCycle = true
		return
	}
	if visited[index] {
		return
	}
	onPath[index] = true
	visited[index] = true
	for _, i2 := range graph[index] {
		traverse(graph, i2)
	}
	list = append(list, index)
	onPath[index] = false
}

func buildGraph(numCourse int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourse)
	for i := 0; i < numCourse; i++ {
		graph[i] = []int{}
	}
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	return graph
}

func main() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	res := findOrder(numCourses, prerequisites)

	fmt.Println(res)

}
