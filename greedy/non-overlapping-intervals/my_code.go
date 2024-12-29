package main

import (
	"fmt"
	"sort"
)

func main() {
	var intervals = [][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 3},
	}

	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	var result int
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			// 别忘记取最小值min
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
			result++
		}
	}

	return result
}
