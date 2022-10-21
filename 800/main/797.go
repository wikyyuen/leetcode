package main

import "fmt"

//797. All Paths From Source to Target
//所有可能的路径

var graphInt [][]int
var result [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	graphInt = graph
	result = [][]int{}
	traverse(0, []int{0})
	return result
}

//solution1
//func traverse(index int, path []int) {
//	if index == len(graphInt)-1 {
//		result = append(result, path)
//		return
//	}
//	for _, i2 := range graphInt[index] {
//		newPath := make([]int, len(path))
//		copy(newPath, path)
//		traverse(i2, append(newPath, i2))
//	}
//	return
//}

//solution2
func traverse(index int, path []int) {
	if index == len(graphInt)-1 {
		newPath := make([]int, len(path))
		copy(newPath, path)
		result = append(result, newPath)
		return
	}
	for _, i2 := range graphInt[index] {
		path = append(path, i2)
		traverse(i2, path)
		path = path[:len(path)-1]
	}
	return
}

func main() {
	graph := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	res := allPathsSourceTarget(graph)
	fmt.Println(res)
}
