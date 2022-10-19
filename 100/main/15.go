package main

import (
	"fmt"
	"reflect"
	"sort"
)

//15. 3Sum

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	lastI := nums[0] + 1
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if nums[i] == lastI {
			continue
		}
		lastI = nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				for nums[right] == nums[right-1] && right-1 > left {
					right--
				}
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				for nums[left] == nums[left+1] && left+1 < right {
					left++
				}
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for nums[left] == nums[left+1] && left+1 < right {
					left++
				}
				left++
				for nums[right] == nums[right-1] && right-1 > left {
					right--
				}
				right--
			}
		}
	}
	return result
}

func inArray(nums [][]int, num []int) bool {
	for i := 0; i < len(nums); i++ {
		if reflect.DeepEqual(nums[i], num) {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}

	result := threeSum(nums)

	fmt.Println(result)
}
