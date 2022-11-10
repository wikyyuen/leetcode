package main

import "fmt"

//39. Combination Sum
//39. 组合总和
var gCandidates []int
var gTarget int
var gGroup []int

func combinationSum(candidates []int, target int) [][]int {
	gCandidates = candidates
	gTarget = target
	gGroup = []int{}
	resList := [][]int{}
	for i := 0; i < len(candidates); i++ {
		traverse2(&resList, i, 0)
	}
	return resList
}

func traverse2(resList *[][]int, index int, sum int) {
	num := gCandidates[index]
	if sum+num > gTarget {
		return
	}
	if sum+num == gTarget {
		temp := make([]int, len(gGroup))
		copy(temp, gGroup)
		temp = append(temp, num)
		*resList = append(*resList, temp)
		return
	}

	for i := index; i < len(gCandidates); i++ {
		gGroup = append(gGroup, num)
		traverse2(resList, i, sum+num)
		gGroup = gGroup[:len(gGroup)-1]
	}
	return
}

func main() {
	//candidates := []int{2, 3, 6, 7}
	//target := 7

	candidates := []int{2, 3, 5}
	target := 8

	res := combinationSum(candidates, target)

	fmt.Println(res)
}
