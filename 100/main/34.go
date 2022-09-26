package main

import "fmt"

//34. Find First and Last Position of Element in Sorted Array

func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, target)
	right := searchRight(nums, target)
	return []int{left, right}
}

func searchLeft(nums []int, target int) int {
	len := len(nums)
	if len == 0 {
		return -1
	}
	if len == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	if len == 2 {
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		} else {
			return -1
		}
	}
	half := len / 2

	if nums[half] >= target {
		return searchLeft(nums[:half+1], target)
	} else {
		new := searchLeft(nums[half:], target)
		if new != -1 {
			return new + half
		}
		return new
	}
}

func searchRight(nums []int, target int) int {
	len := len(nums)
	if len == 0 {
		return -1
	}
	if len == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	if len == 2 {
		if nums[1] == target {
			return 1
		} else if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	half := len / 2

	if nums[half] > target {
		return searchRight(nums[:half+1], target)
	} else {
		new := searchRight(nums[half:], target)
		if new != -1 {
			return new + half
		}
		return new
	}
}

func main() {
	//nums := []int{5, 6, 7, 7, 8, 8, 8, 10}
	nums := []int{}
	target := 1

	result := searchRange(nums, target)
	fmt.Println(result)
}
