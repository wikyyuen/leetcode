package main

import (
	"fmt"
	"sort"
)

//354. Russian Doll Envelopes
//354. 俄罗斯套娃信封问题

func maxEnvelopes(envelopes [][]int) int {
	//排序
	sort.SliceStable(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	//dp
	//dpList := make([]int, n)
	//maxNum := 1
	//dpList[0] = 1
	//for i := 1; i < n; i++ {
	//	dpList[i] = 1
	//	for j := 0; j < i; j++ {
	//		if envelopes[j][1] < envelopes[i][1] {
	//			dpList[i] = max(dpList[j]+1, dpList[i])
	//
	//		}
	//	}
	//	maxNum = max(maxNum, dpList[i])
	//}

	//二分
	minHeight := []int{0}
	for _, envelope := range envelopes {
		// 寻找 minHeight 中第一个大于等于 envelope[1] 的下标 k ，
		// 则说明以当前信封为最外层时，最多能嵌套 k 层
		k := sort.Search(len(minHeight), func(i int) bool { return minHeight[i] >= envelope[1] })
		if len(minHeight) == k {
			// 当前信封是第一个嵌套 k 层的信封，所以直接放入 minHeight
			minHeight = append(minHeight, envelope[1])
		} else {
			// 此时存在嵌套 k 个信封的情况，
			// 因为前面二分寻找的是第一个大于等于 envelope[1] 的下标，
			// 所以 minHeight[k] >= envelope[1] ，可以直接更新为 envelope[1]
			minHeight[k] = envelope[1]
		}
	}

	return len(minHeight) - 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	envelopes := [][]int{{15, 22}, {8, 34}, {44, 40}, {9, 17}, {43, 23}, {4, 7}, {20, 8}, {30, 46}, {39, 36}, {45, 14}, {24, 19}, {24, 36}, {31, 34}, {32, 19}, {29, 13}, {16, 48}, {8, 36}, {44, 2}, {11, 5}, {2, 50}, {29, 6}, {18, 38}, {15, 49}, {22, 37}, {6, 20}, {25, 11}, {1, 50}, {19, 40}, {45, 35}, {37, 21}, {4, 29}, {40, 5}, {4, 49}, {1, 45}, {14, 32}, {14, 37}, {23, 22}, {31, 21}, {2, 36}, {43, 4}, {21, 32}, {41, 2}, {44, 32}, {36, 20}, {22, 45}, {3, 41}, {44, 29}, {29, 33}, {42, 2}, {38, 17}, {43, 26}, {30, 15}, {28, 12}, {33, 30}, {49, 7}, {8, 14}, {1, 9}, {41, 25}, {7, 15}, {26, 32}, {11, 33}, {12, 45}, {33, 7}, {16, 34}, {39, 1}, {20, 49}, {50, 45}, {14, 29}, {50, 41}, {1, 45}, {14, 43}, {49, 20}, {41, 37}, {43, 22}, {45, 19}, {20, 21}, {28, 19}, {2, 1}, {7, 49}, {3, 3}, {49, 48}, {34, 35}, {10, 2}}

	res := maxEnvelopes(envelopes)

	fmt.Println(res)
}
