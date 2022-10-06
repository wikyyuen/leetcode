package main

import "fmt"

//283. Move Zeroes
//solution1
//func moveZeroes(nums []int) []int {
//	n := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == 0 {
//			nums = append(nums[:i], nums[i+1:]...)
//			i--
//			n++
//		}
//	}
//	addArr := make([]int, n)
//	nums = append(nums, addArr...)
//	return nums
//}
//solution2
func moveZeroes(nums []int) {
	index := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)

	fmt.Println(nums)
}
