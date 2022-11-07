package main

import (
	"fmt"
	"sort"
)

//322. Coin Change
//322. 零钱兑换

var dpList []int
var gCoins []int

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dpList = make([]int, amount+1)
	gCoins = coins
	sort.Ints(gCoins)
	res := dp(amount)
	if res == 0 {
		res = -1
	}
	return res
}

func dp(amount int) int {
	//if amount <= 0 {
	//	return 0
	//}
	if dpList[amount] != 0 {
		return dpList[amount]
	}
	min := 99999
	for i := 0; i < len(gCoins); i++ {
		if gCoins[i] > amount {
			break
		}
		if gCoins[i] == amount {
			return 1
		}
		if dp(amount-gCoins[i]) < min && dp(amount-gCoins[i]) > 0 {
			min = dp(amount - gCoins[i])
		}
	}
	if min == 99999 {
		min = -1
	} else {
		min += 1
	}
	dpList[amount] = min
	return min
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11

	result := coinChange(coins, amount)

	fmt.Println(result)
}
