package main

import (
	"fmt"
	"strconv"
)

var result []string
var path []string

func main() {
	s := "aab"
	restoreIpAddresses(s)
	fmt.Println(result)
}

func restoreIpAddresses(s string) []string {
	result = []string{}
	path = []string{}
	backtracking(s, path)
	return result
}

func backtracking(s string, path []string) {
	if len(s) == 0 && len(path) == 4 {
		subResult := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
		result = append(result, subResult)
		return
	}

	for i := 0; i < len(s); i++ {
		sNum := -1
		subS := s[:i+1]
		//fmt.Println(subS)
		if len(subS) > 1 && subS[0] == '0' {
			continue
		}
		sNum, _ = strconv.Atoi(subS)
		if sNum >= 0 && sNum <= 255 {
			path = append(path, subS)
			//fmt.Println(path)
			//fmt.Println(s[i+1:])
			if len(path) > 4 {
				break
			}
			backtracking(s[i+1:], path)
			path = path[:len(path)-1]
		} else {
			continue
		}
	}
}
