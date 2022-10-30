package main

import (
	"fmt"
)

//46. Permutations
//46. 全排列

func permute(nums []int) [][]int {
	res := [][]int{}
	for _, i2 := range nums {
		res = traverse(res, i2)
	}
	return res
}

func traverse(nums [][]int, n int) [][]int {
	if len(nums) == 0 {
		return [][]int{{n}}
	}
	res := [][]int{}
	for _, item := range nums {
		temp := make([]int, len(item)+1)
		copy(temp, append([]int{n}, item...))
		res = append(res, temp)
		temp = make([]int, len(item)+1)
		copy(temp, append(item, n))
		res = append(res, temp)

		for i := 1; i < len(item); i++ {
			res = append(res, addNum(item, n, i))
		}
	}

	return res
}

func addNum(nums []int, n int, index int) []int {
	temp := make([]int, len(nums)+1)
	copy(temp, append(nums[:index], append([]int{n}, nums[index:]...)...))
	return temp
}

func main() {
	//nums := []int{1, 2, 3}
	nums := []int{5, 4, 6, 2}

	result := permute(nums)

	fmt.Println(result)
}
