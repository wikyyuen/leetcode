package main

//231. Power of Two
//231. 2 的幂

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return (n & (n - 1)) == 0
}
