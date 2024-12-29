package main

import (
	"fmt"
	"sort"
)

func main() {
	var s = "ababcbacadefegdehijhklij"

	fmt.Println(partitionLabels(s))
}

func partitionLabels(s string) []int {
	var result []int
	var sMap = make(map[byte][]int)
	var sArr [][]int

	for i := range s {
		sMap[s[i]] = append(sMap[s[i]], i)
	}

	for i := range sMap {
		arr := []int{sMap[i][0], sMap[i][len(sMap[i])-1]}
		sArr = append(sArr, arr)
	}

	sort.Slice(sArr, func(i, j int) bool {
		if sArr[i][1] == sArr[j][1] {
			return sArr[i][1] < sArr[j][1]
		}
		return sArr[i][0] < sArr[j][0]
	})

	fmt.Println(sArr)

	for i := 1; i < len(sArr); i++ {
		if sArr[i][0] > sArr[i-1][1] {
			//fmt.Printf("sArr[%d][1]: %d, sArr[%d][0]: %d\n", i-1, sArr[i-1][1], i-1, sArr[i-1][0])
			result = append(result, sArr[i-1][1]-sArr[i-1][0]+1)
		} else {
			sArr[i][1] = max(sArr[i][1], sArr[i-1][1])
			sArr[i][0] = min(sArr[i][0], sArr[i-1][0])

		}
		// fmt.Println(sArr[i][1])
		// fmt.Println(sArr[i][0])
		if i == len(sArr)-1 {
			result = append(result, sArr[i][1]-sArr[i][0]+1)
		}
	}

	return result
}
