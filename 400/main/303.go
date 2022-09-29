package main

import "fmt"

//303. Range Sum Query - Immutable

type NumArray struct {
	PreSum []int
}

func Constructor2(nums []int) NumArray {
	preLen := len(nums) + 1
	preSum := make([]int, preLen)

	preSum[0] = 0
	for i := 1; i < preLen; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{
		PreSum: preSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.PreSum[right+1] - this.PreSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {
	numbers := []int{-2, 0, 3, -5, 2, -1}

	numArray := Constructor2(numbers)

	result := numArray.SumRange(0, 3)

	fmt.Println(result)
}
