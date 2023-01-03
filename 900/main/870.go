package main

import (
	"fmt"
	"sort"
)

//870. Advantage Shuffle
//870. 优势洗牌

func advantageCount(nums1 []int, nums2 []int) []int {
	var newNums2 [][]int
	var res []int
	n := len(nums1)

	sort.Ints(nums1)

	newNums2 = make([][]int, n)

	nums2Copy := make([]int, n)
	res = make([]int, n)

	copy(nums2Copy, nums2)
	sort.Ints(nums2Copy)

	for i := 0; i < n; i++ {
		newNums2[i] = make([]int, 2)
		newNums2[i][0] = i
		newNums2[i][1] = nums2[i]
	}
	sort.Slice(newNums2, func(i int, j int) bool {
		return newNums2[i][1] > newNums2[j][1]
	})
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums1[right] > newNums2[i][1] {
			res[newNums2[i][0]] = nums1[right]
			right--
		} else {
			res[newNums2[i][0]] = nums1[left]
			left++
		}
	}
	return res
}

func main() {
	nums1 := []int{718967141, 189971378, 341560426, 23521218, 339517772}
	nums2 := []int{967482459, 978798455, 744530040, 3454610, 940238504}

	result := advantageCount(nums1, nums2)

	fmt.Println(result)
}
