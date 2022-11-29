package main

import "fmt"

//268. Missing Number
//268. 丢失的数字

func missingNumber(nums []int) int {
	//n := len(nums)
	//res := 0 ^ n
	//for i := 0; i < n; i++ {
	//	res ^= nums[i] ^ i
	//}
	sum := len(nums) * (len(nums) + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	res := missingNumber(nums)

	fmt.Println(res)
}
