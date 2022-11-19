package main

import "fmt"

//136. Single Number
//136. 只出现一次的数字

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func main() {
	nums := []int{4, 1, 2, 1, 2}

	res := singleNumber(nums)

	fmt.Println(res)
}
