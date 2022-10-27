package main

import (
	"fmt"
	"sort"
)

//698. Partition to K Equal Sum Subsets
//698. 划分为k个相等的子集
var perSum int
var count int
var gNums []int
var cup []int

func canPartitionKSubsets(nums []int, k int) bool {
	sort.Ints(nums)
	//反射并求和
	total := 0
	for _, num := range nums {
		total += num
	}
	gNums = nums
	if total%k != 0 {
		return false
	}
	count = k
	perSum = total / k
	if nums[len(nums)-1] > perSum {
		return false
	}
	cup = make([]int, k)
	for i := 0; i < k; i++ {
		cup[i] = perSum
	}
	return traverse(len(nums) - 1)
}

func traverse(idx int) bool {
	if idx < 0 {
		return true
	}
	for i := range cup {
		if cup[i] == gNums[idx] || cup[i]-gNums[idx] >= gNums[0] {
			cup[i] -= gNums[idx]
			if traverse(idx - 1) {
				return true
			}
			cup[i] += gNums[idx]
		}
		if cup[i] == perSum {
			break
		}
	}
	return false
}

//func sum(nums []int) {
//	total := 0
//	for i, num := range nums {
//		total += num
//	}
//	return total
//}

func main() {
	//nums := []int{4, 3, 2, 3, 5, 2, 1}
	//k := 4
	//nums := []int{1, 2, 3, 4}
	//k := 3
	nums := []int{3522, 181, 521, 515, 304, 123, 2512, 312, 922, 407, 146, 1932, 4037, 2646, 3871, 269}
	k := 5
	//nums := []int{4, 5, 9, 3, 10, 2, 10, 7, 10, 8, 5, 9, 4, 6, 4, 9}
	//nums := []int{4, 3, 2, 3, 5, 2, 1}
	//k := 4
	result := canPartitionKSubsets(nums, k)

	fmt.Println(result)
}
