package main

//191. Number of 1 Bits
//191. 位1的个数

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}

func main() {

}
