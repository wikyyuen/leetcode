package main

//40. Combination Sum II
//40. 组合总和 II

import (
	"fmt"
	"sort"
)

var gCandidates []int
var gTarget int
var gGroup []int

func combinationSum2(candidates []int, target int) [][]int {
	gCandidates = candidates
	sort.Ints(gCandidates)
	gTarget = target
	gGroup = []int{}
	res := [][]int{}
	traverse4(&res, 0, 0)
	return res
}

func traverse4(res *[][]int, start int, sum int) {
	for i := start; i < len(gCandidates); i++ {
		if i > start && gCandidates[i] == gCandidates[i-1] {
			continue
		}
		num := gCandidates[i]
		if num+sum > gTarget {
			return
		}
		if num+sum == gTarget {
			temp := make([]int, len(gGroup))
			copy(temp, gGroup)
			temp = append(temp, num)
			*res = append(*res, temp)
			return
		}

		gGroup = append(gGroup, num)
		traverse4(res, i+1, sum+num)
		gGroup = gGroup[:len(gGroup)-1]
	}
	return
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	res := combinationSum2(candidates, target)

	fmt.Println(res)
}
