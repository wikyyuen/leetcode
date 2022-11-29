package main

import (
	"fmt"
)

//55. 跳跃游戏
//55. Jump Game

func canJump(nums []int) bool {
	f := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if f < i+nums[i] {
			f = i + nums[i]
		}
		if f <= i {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	res := canJump(nums)
	fmt.Println(res)
}
