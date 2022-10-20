package main

import (
	"fmt"
	"sort"
)

//16. 3Sum Closest

//solution 1
//func threeSumClosest(nums []int, target int) int {
//	if len(nums) < 3 {
//		return 0
//	}
//	max := nums[0] + nums[1] + nums[2]
//	temp := max
//	if len(nums) > 3 {
//		//排序
//		//sort.Ints(nums)
//		for i := 0; i < len(nums)-2; i++ {
//			for j := i + 1; j < len(nums)-1; j++ {
//				for k := j + 1; k < len(nums); k++ {
//					temp = nums[i] + nums[j] + nums[k]
//					if abs(temp-target) < abs(max-target) {
//						max = temp
//					}
//				}
//			}
//		}
//	}
//	return max
//}
//solution2
var closest int
var diff int

func threeSumClosest(nums []int, target int) int {
	closest = nums[0] + nums[1] + nums[2]
	diff = 100000
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] > target {
				if check(nums[i]+nums[left]+nums[right], target) {
					return closest
				}
				for nums[right] == nums[right-1] && right-1 > left {
					right--
				}
				right--
			} else if nums[i]+nums[left]+nums[right] < target {
				if check(nums[i]+nums[left]+nums[right], target) {
					return closest
				}
				for nums[left] == nums[left+1] && left+1 < right {
					left++
				}
				left++
			} else {
				return target
			}
		}
	}
	return closest
}

func check(n int, target int) bool {
	if abs(n-target) < diff {
		diff = abs(n - target)
		closest = n
	}
	if diff == 0 {
		return true
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	nums := []int{1, 1, 1, 0}
	target := -100

	result := threeSumClosest(nums, target)

	fmt.Println(result)
}
