package main

import "fmt"

//496. Next Greater Element I
//496. 下一个更大元素 I

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	gater := nextGreaterElement2(nums2)
	gaterMap := map[int]int{}
	for i := 0; i < len(gater); i++ {
		gaterMap[nums2[i]] = gater[i]
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = gaterMap[nums1[i]]
	}
	return res
}

func nextGreaterElement2(nums1 []int) []int {
	res := make([]int, len(nums1))
	stack := []int{}
	for i := len(nums1) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[0] < nums1[i] {
			stack = stack[1:]
		}
		if len(stack) == 0 {
			res[i] = -1
			stack = append(stack, nums1[i])
		} else {
			res[i] = stack[0]
			stack = append([]int{nums1[i]}, stack...)
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	result := nextGreaterElement(nums1, nums2)
	fmt.Println(result)
}
