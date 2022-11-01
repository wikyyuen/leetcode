package main

import "fmt"

//300. Longest Increasing Subsequence
//300. 最长递增子序列

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	max := 0
	for i := 0; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func main() {
	//nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums := []int{7, 7, 7, 7, 7, 7, 7}
	nums := []int{0, 1, 0, 3, 2, 3}

	result := lengthOfLIS(nums)

	fmt.Println(result)
}
