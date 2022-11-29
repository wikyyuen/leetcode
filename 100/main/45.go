package main

import "fmt"

//45. Jump Game II
//45. 跳跃游戏 II

func jump(nums []int) int {
	//贪心
	j, e, s := 0, 0, 0
	for e < len(nums)-1 {
		j++
		farthest := e
		for i := s; i <= e; i++ {
			if farthest < i+nums[i] {
				farthest = i + nums[i]
			}
		}
		s, e = e+1, farthest
	}
	return j
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{2, 3, 0, 1, 4}
	res := jump(nums)
	fmt.Println(res)

}
