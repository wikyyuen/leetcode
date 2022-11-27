package main

import (
	"fmt"
	"sort"
)

//1024. Video Stitching
//1024. 视频拼接

func videoStitching(clips [][]int, time int) int {
	//先排序
	sort.SliceStable(clips, func(i int, j int) bool {
		if clips[j][0] == clips[i][0] {
			return clips[j][1] < clips[i][1]
		} else {
			return clips[j][0] > clips[i][0]
		}
	})
	if clips[0][0] != 0 {
		return -1
	}

	curEnd, nextEnd, i, res := 0, 0, 0, 0
	n := len(clips)
	for i < n && clips[i][0] <= curEnd {
		for i < n && clips[i][0] <= curEnd {
			if clips[i][1] > nextEnd {
				nextEnd = clips[i][1]
			}
			i++
		}
		res++
		curEnd = nextEnd
		if curEnd >= time {
			return res
		}
	}
	return -1
}

func main() {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	time := 10

	res := videoStitching(clips, time)

	fmt.Println(res)
}
