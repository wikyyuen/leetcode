package main

import "fmt"

//34. Find First and Last Position of Element in Sorted Array

func searchRange(nums []int, target int) []int {
	len := len(nums)
	half := len / 2
	return []int{half, -1}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	result := searchRange(nums, target)
	fmt.Println(result)
}
