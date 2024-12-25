package main

import "fmt"

var result []string
var path []int

func main() {
	digits := "23"
	letterCombinations(digits)
	fmt.Println(result)
}

func letterCombinations(digits string) []string {
	result = []string{}
	path = []int{}
	backtracking(digits, "", 0)
	return result
}

func backtracking(digits string, path string, startIndex int) {
	if len(digits) == 0 {
		return
	}
	if len(digits) == len(path) {
		result = append(result, path)
		return
	}

	for i := startIndex; i < len(digits); i++ {
		var sOfDigits []string
		switch digits[i] {
		case '2':
			sOfDigits = []string{"a", "b", "c"}
		case '3':
			sOfDigits = []string{"d", "e", "f"}
		case '4':
			sOfDigits = []string{"g", "h", "i"}
		case '5':
			sOfDigits = []string{"j", "k", "l"}
		case '6':
			sOfDigits = []string{"m", "n", "o"}
		case '7':
			sOfDigits = []string{"p", "q", "r", "s"}
		case '8':
			sOfDigits = []string{"t", "u", "v"}
		case '9':
			sOfDigits = []string{"w", "x", "y", "z"}
		}

		for j := range sOfDigits {
			path = path + sOfDigits[j]
			backtracking(digits, path, i+1)
			path = path[:len(path)-1]
		}
	}
}
