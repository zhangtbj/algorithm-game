package main

import "fmt"

var result [][]string
var path []string

func main() {
	s := "aab"
	partition(s)
	fmt.Println(result)
}

func partition(s string) [][]string {
	result = [][]string{}
	path = []string{}
	backtracking(s, path)
	return result
}

func backtracking(s string, path []string) {
	//fmt.Println(path)
	if len(path) > 0 && !checkPartition(path[len(path)-1]) {
		return
	}

	if len(s) == 0 {
		tmp := make([]string, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := 0; i < len(s); i++ {
		subS := s[:i+1]
		s := s[i+1:]
		fmt.Printf("subS: %s, restS: %s\n", subS, s)
		path = append(path, subS)
		fmt.Println(path)
		backtracking(s, path)
		path = path[:len(path)-1]
		s = subS + s
	}
}

func checkPartition(s string) bool {
	start, end := 0, len(s)-1
	for start <= end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
