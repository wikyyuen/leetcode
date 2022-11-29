package main

import "fmt"

//45. Jump Game II
//45. 跳跃游戏 II

var gNums []int
var dpList []int

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	gNums = nums
	dpList = make([]int, len(nums))
	return dp(0)
}

func dp(index int) int {
	if gNums[index] == 0 {
		return len(gNums)
	}
	if gNums[index]+index >= len(gNums)-1 {
		return 1
	}

	if dpList[index] != 0 {
		return dpList[index]
	}
	min := len(gNums)
	for i := index + 1; i <= gNums[index]+index; i++ {
		temp := dp(i)
		if min > temp {
			min = temp
		}
	}
	dpList[index] = min + 1
	return dpList[index]
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{2, 3, 0, 1, 4}
	res := jump(nums)
	fmt.Println(res)

}
