package main

import (
	"fmt"
	"sort"
)

func main() {
	var points = [][]int{
		{10, 16}, {2, 8}, {1, 6}, {7, 12},
	}

	fmt.Println(findMinArrowShots(points))
}

func findMinArrowShots(points [][]int) int {
	var result int

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			result++
		} else {
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}

	return result + 1
}
