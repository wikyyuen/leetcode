package main

import (
	"fmt"
	"math/rand"
)

//380. Insert Delete GetRandom O(1)

type RandomizedSet struct {
	KeyMap map[int]int
	Nums   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Nums:   []int{},
		KeyMap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.KeyMap[val]; ok {
		return false
	}
	this.KeyMap[val] = len(this.Nums)
	this.Nums = append(this.Nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if key, ok := this.KeyMap[val]; ok {
		length := len(this.Nums)
		lastValue := this.Nums[length-1]
		this.Nums[key] = lastValue
		this.KeyMap[lastValue] = key
		delete(this.KeyMap, val)
		this.Nums = this.Nums[:length-1]
		return true
	} else {
		return false
	}

}

func (this *RandomizedSet) GetRandom() int {
	return this.Nums[rand.Int31n(int32(len(this.Nums)))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
func main() {
	obj := Constructor()

	fmt.Println(obj.Insert(0))
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.GetRandom())

}
