package main

import "fmt"

//990. Satisfiability of Equality Equations
//等式方程的可满足性

var parent []int
var size []int

func equationsPossible(equations []string) bool {
	parent = make([]int, 26)
	size = make([]int, 26)
	equations2 := []string{}
	//初始化父节点
	for i, _ := range parent {
		parent[i] = i
		size[i] = 1
	}

	for _, equation := range equations {
		tempByte := []byte(equation)
		if string(tempByte[1]) == "=" {
			if tempByte[0] != tempByte[3] {
				union(int(tempByte[0])-97, int(tempByte[3])-97)
			}
		} else {
			equations2 = append(equations2, equation)
		}
	}
	for _, equation := range equations2 {
		tempByte := []byte(equation)
		if connected(int(tempByte[0])-97, int(tempByte[3])-97) {
			return false
		}
	}
	return true
}

func union(a int, b int) {
	rootA := find(a)
	rootB := find(b)

	if rootB != rootA {
		//合并
		if size[rootA] < size[rootB] {
			parent[rootA] = rootB
			size[rootB] += size[rootA]
		} else {
			parent[rootB] = rootA
			size[rootA] += size[rootB]
		}
	}
}

func connected(a int, b int) bool {
	rootA := find(a)
	rootB := find(b)

	return rootB == rootA
}

//查父节点并压缩树
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func main() {
	//equations := []string{"a==b", "b!=a"}
	//equations := []string{"b==a", "a==b"}
	equations := []string{"c==c", "b==d", "x!=z"}

	result := equationsPossible(equations)

	fmt.Println(result)

}
