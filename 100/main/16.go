package main

import (
	"fmt"
)

//16. 3Sum Closest

//solution 1
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	max := nums[0] + nums[1] + nums[2]
	temp := max
	if len(nums) > 3 {
		//排序
		//sort.Ints(nums)
		for i := 0; i < len(nums)-2; i++ {
			for j := i + 1; j < len(nums)-1; j++ {
				for k := j + 1; k < len(nums); k++ {
					temp = nums[i] + nums[j] + nums[k]
					if abs(temp-target) < abs(max-target) {
						max = temp
					}
				}
			}
		}
	}
	return max
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	nums := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	target := -2

	result := threeSumClosest(nums, target)

	fmt.Println(result)
}
