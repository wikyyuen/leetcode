package main

import (
	"fmt"
	"strconv"
)

//659. 分割数组为连续子序列
//659. Split Array into Consecutive Subsequences

func isPossible(nums []int) bool {
	freNums := make(map[string]int)
	need := make(map[string]int)
	for _, item := range nums {
		v := strconv.Itoa(item)
		freNums[v]++
	}
	for _, item := range nums {
		v := strconv.Itoa(item)
		v1 := strconv.Itoa(item + 1)
		v2 := strconv.Itoa(item + 2)
		v3 := strconv.Itoa(item + 3)
		if freNums[v] == 0 {
			continue
		}
		if need[v] > 0 {
			freNums[v]--
			need[v]--
			need[v1]++
		} else if freNums[v] > 0 && freNums[v1] > 0 && freNums[v2] > 0 {
			freNums[v]--
			freNums[v1]--
			freNums[v2]--
			need[v3]++
		} else {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1, 2, 3, 3, 4, 5}
	result := isPossible(nums)
	fmt.Println(result)
}
