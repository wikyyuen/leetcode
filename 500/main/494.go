package main

import (
	"fmt"
	"strconv"
)

//494. Target Sum
//494. 目标和
var dpList map[string]int

func findTargetSumWays(nums []int, target int) int {
	dpList = make(map[string]int)
	if len(nums) == 0 {
		return 0
	}
	return dp(nums, 0, target)
}

func dp(nums []int, index int, target int) int {
	if index == len(nums)-1 {
		res := 0
		if nums[index]-target == 0 {
			res++
		}
		if nums[index]+target == 0 {
			res++
		}
		return res
	}
	key := strconv.Itoa(index) + "," + strconv.Itoa(target)
	value, ok := dpList[key]
	if ok {
		return value
	}
	dpList[key] = dp(nums, index+1, target-nums[index]) + dp(nums, index+1, target+nums[index])
	return dpList[key]
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3

	res := findTargetSumWays(nums, target)

	fmt.Println(res)
}
