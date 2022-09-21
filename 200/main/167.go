package main

import (
	"fmt"
)

//167. Two Sum II - Input Array Is Sorted
//solution1
//func twoSum(numbers []int, target int) []int {
//	slow := 0
//	fast := 1
//	len := len(numbers)
//
//	for fast < len {
//		if numbers[slow]+numbers[fast] == target {
//			return []int{slow + 1, fast + 1}
//		}
//		slow++
//		if slow == fast || numbers[slow]+numbers[fast] > target {
//			slow = 0
//			fast++
//		}
//	}
//	return []int{}
//}

//func twoSum(numbers []int, target int) []int {
//	slow, fast := 0, len(numbers)-1
//
//	for slow < fast {
//		sum := numbers[slow] + numbers[fast]
//		if sum == target {
//			return []int{slow + 1, fast + 1}
//		}
//		if sum < target {
//			slow++
//		} else {
//			fast--
//		}
//	}
//	return []int{-1, -1}
//}

func twoSum(numbers []int, target int) []int {
	res := make(map[int]int, len(numbers))

	for i := range numbers {
		l := target - numbers[i]
		if j, ok := res[l]; ok {
			return []int{j + 1, i + 1}
		}
		res[numbers[i]] = i
	}
	return []int{}
}

func main() {

	//numbers := []int{-1, 0}
	//target := -1

	//numbers := []int{2, 7, 11, 15}
	//target := 9

	numbers := []int{12, 13, 23, 28, 43, 44, 59, 60, 61, 68, 70, 86, 88, 92, 124, 125, 136, 168, 173, 173, 180, 199, 212, 221, 227, 230, 277, 282, 306, 314, 316, 321, 325, 328, 336, 337, 363, 365, 368, 370, 370, 371, 375, 384, 387, 394, 400, 404, 414, 422, 422, 427, 430, 435, 457, 493, 506, 527, 531, 538, 541, 546, 568, 583, 585, 587, 650, 652, 677, 691, 730, 737, 740, 751, 755, 764, 778, 783, 785, 789, 794, 803, 809, 815, 847, 858, 863, 863, 874, 887, 896, 916, 920, 926, 927, 930, 933, 957, 981, 997}
	target := 542

	result := twoSum(numbers, target)

	fmt.Println(result)
}
