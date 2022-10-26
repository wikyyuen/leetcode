package main

import "fmt"

//216. Combination Sum III
//216. 组合总和 III

var (
	track []int
	res2  [][]int
)

func combinationSum3(k int, n int) [][]int {
	res2 = [][]int{}
	track = []int{}
	traverse(1, n, 0, k)
	return res2
}

func traverse(start int, n int, total int, k int) {
	for i := start; i < 10; i++ {
		track = append(track, i)
		if len(track) < k {
			traverse(i+1, n, total+i, k)
		}
		if total+i == n && len(track) == k {
			temp := make([]int, len(track))
			copy(temp, track)
			res2 = append(res2, temp)
			track = track[:len(track)-1]
			break
		}

		track = track[:len(track)-1]
		if total+i > n {
			break
		}
	}
}

func main() {
	k := 4
	n := 1

	result := combinationSum3(k, n)

	fmt.Println(result)
}
