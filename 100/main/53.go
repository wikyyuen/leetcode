package main

import "fmt"

//53. Maximum Subarray
//53. 最大子数组和

func maxSubArray(nums []int) int {
	left, right, windowSum := 0, 0, 0
	maxSum := -9999999
	for right < len(nums) {
		windowSum += nums[right]
		right++

		if windowSum > maxSum {
			maxSum = windowSum
		}
		for windowSum < 0 {
			windowSum -= nums[left]
			left++
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums)
	fmt.Println(result)
}
