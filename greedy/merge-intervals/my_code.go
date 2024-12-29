package main

import (
	"fmt"
	"sort"
)

func main() {
	var intervals = [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}

	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	l := len(intervals)

	if l == 1 {
		return intervals
	}

	var result [][]int
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	for i := 1; i < l; i++ {
		if intervals[i][0] > intervals[i-1][1] {
			result = append(result, intervals[i-1])
		} else {
			intervals[i][0] = min(intervals[i-1][0], intervals[i][0])
			intervals[i][1] = max(intervals[i-1][1], intervals[i][1])
		}

		if i == l-1 {
			result = append(result, intervals[i])
		}
	}

	return result
}
