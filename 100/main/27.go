package main

import "fmt"

//27. Remove Element

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	return index
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	result := removeElement(nums, val)
	fmt.Println(result)
	fmt.Println(nums)
}
