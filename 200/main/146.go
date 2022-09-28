package main

import "fmt"

//146. LRU Cache

type LRUCache struct {
	KeyMap    map[int]*DoubleList
	Capacity  int
	LastNode  *DoubleList
	FirstNode *DoubleList
}

type DoubleList struct {
	Next *DoubleList
	Pre  *DoubleList
	Key  int
	Val  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		KeyMap:   make(map[int]*DoubleList),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if doubleList, ok := this.KeyMap[key]; ok {
		this.RefreshNodePos(doubleList)
		return doubleList.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if doubleList, ok := this.KeyMap[key]; ok {
		this.RefreshNodePos(doubleList)
		doubleList.Val = value
	} else {
		if len(this.KeyMap) == this.Capacity {
			//full
			delete(this.KeyMap, this.LastNode.Key)
			if this.LastNode.Pre != nil {
				this.LastNode = this.LastNode.Pre
			}
			this.LastNode.Next = nil
		}
		doubleList = &DoubleList{
			Key: key,
			Val: value,
		}
		doubleList.Pre = nil
		if this.FirstNode != nil {
			doubleList.Next = this.FirstNode
			this.FirstNode.Pre = doubleList
		}
		this.FirstNode = doubleList
		if len(this.KeyMap) == 0 {
			this.LastNode = doubleList
		}
		this.KeyMap[key] = doubleList
	}
}

func (this *LRUCache) RefreshNodePos(doubleList *DoubleList) {
	if len(this.KeyMap) > 1 && this.FirstNode != doubleList {
		//移除节点
		if doubleList.Next != nil {
			doubleList.Next.Pre = doubleList.Pre
		}
		if doubleList.Pre != nil {
			doubleList.Pre.Next = doubleList.Next
		}
		if this.LastNode == doubleList {
			this.LastNode = doubleList.Pre
		}
		doubleList.Pre = nil
		doubleList.Next = this.FirstNode
		this.FirstNode.Pre = doubleList
		this.FirstNode = doubleList
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	capacity := 2
	//capacity := 1
	obj := Constructor(capacity)
	//obj.Put(1, 1)
	//obj.Put(2, 2)
	//fmt.Println(obj.Get(1))
	//obj.Put(3, 3)
	//fmt.Println(obj.Get(2))
	//obj.Put(4, 4)
	//fmt.Println(obj.Get(1))
	//fmt.Println(obj.Get(3))
	//fmt.Println(obj.Get(4))

	//fmt.Println(obj.Get(2))
	//obj.Put(2, 6)
	//fmt.Println(obj.Get(1))
	//obj.Put(1, 5)
	//obj.Put(1, 2)
	//fmt.Println(obj.Get(1))
	//fmt.Println(obj.Get(2))

	//obj.Put(2, 1)
	//obj.Put(2, 2)
	//fmt.Println(obj.Get(2))
	//obj.Put(1, 1)
	//obj.Put(4, 1)
	//fmt.Println(obj.Get(2))

	//obj.Put(1, 1)
	//obj.Put(2, 2)
	//fmt.Println(obj.Get(1))
	//obj.Put(3, 3)
	//fmt.Println(obj.Get(2))
	//obj.Put(4, 4)
	//fmt.Println(obj.Get(1))
	//fmt.Println(obj.Get(3))
	//fmt.Println(obj.Get(4))

	obj.Put(2, 1)
	obj.Put(3, 2)
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(2))
	obj.Put(4, 3)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

}
