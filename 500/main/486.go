package main

import "fmt"

//486. Predict the Winner
//486. 预测赢家

var dpList [][][]int
var gNums []int

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	gNums = nums
	dpList = make([][][]int, n)
	for i := 0; i < n; i++ {
		dpList[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dpList[i][j] = make([]int, 2)
			dpList[i][j][0] = -1
			dpList[i][j][1] = -1
		}
	}
	if dp(0, n-1, 0)-dp(0, n-1, 1) >= 0 {
		return true
	}
	return false
}

func dp(i int, j int, first int) int {
	if j <= i {
		if first == 0 {
			return gNums[i]
		} else {
			return 0
		}
	}
	if dpList[i][j][first] != -1 {
		return dpList[i][j][first]
	}
	max := gNums[i] + dp(i+1, j, 1)
	if max < gNums[j]+dp(i, j-1, 1) {
		if first == 0 {
			dpList[i][j][first] = gNums[j] + dp(i, j-1, 1)
		} else {
			dpList[i][j][first] = dp(i, j-1, 0)
		}
		return dpList[i][j][first]
	}
	if first == 0 {
		dpList[i][j][first] = max
	} else {
		dpList[i][j][first] = dp(i+1, j, 0)
	}
	return dpList[i][j][first]
}

func main() {
	//nums := []int{1, 5, 2}
	nums := []int{1, 5, 233, 7}
	result := PredictTheWinner(nums)
	fmt.Println(result)
}
