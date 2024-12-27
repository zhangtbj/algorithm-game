package main

import (
	"fmt"
	"math"
	"sort"
)

const (
	TARGET                          = 24
	EPSILON                         = 1e-6
	ADD, MULTIPLY, SUBTRACT, DIVIDE = 0, 1, 2, 3
)

var result []string
var path string
var used []bool

func main() {
	cards := []int{4, 1, 8, 7}
	result = []string{}
	path = ""
	used = make([]bool, len(cards))
	sort.Ints(cards)

	judgePoint24(cards)
	//fmt.Println(result)
}

func judgePoint24(cards []int) bool {
	floatNums := make([]float64, len(cards))
	for i := range floatNums {
		floatNums[i] = float64(cards[i])
	}

	backtracking(floatNums)
	fmt.Printf("\nIs Vaild Cards?: %t\n", len(result) > 0)
	fmt.Println(cards)

	return len(result) > 0
}

func backtracking(cards []float64) {
	if len(cards) == 0 {
		return
	}

	if len(cards) == 1 {
		if math.Abs(cards[0]-TARGET) < EPSILON {
			fmt.Println(path)
			fmt.Printf("=%.1f\n", cards[0])

			if len(result) == 0 || (len(result) > 0 && result[len(result)-1] != path) {
				result = append(result, path)
			}
			return
		}
	}

	size := len(cards)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i != j {
				var list2 []float64
				for k := 0; k < size; k++ {
					if k != i && k != j {
						list2 = append(list2, cards[k])
					}
				}
				for k := 0; k < 4; k++ {
					// 加法和乘法都满足交换律，因此如果选择的运算操作是加法或乘法，则对于选出的 2 个数字不需要考虑不同的顺序，在遇到第二种顺序时可以不进行运算，直接跳过。
					if k < 2 && i < j {
						continue
					}
					oldPath := path
					switch k {
					case ADD:
						list2 = append(list2, cards[i]+cards[j])
						path += fmt.Sprintf("%.1f+%.1f,", cards[i], cards[j])
					case MULTIPLY:
						list2 = append(list2, cards[i]*cards[j])
						path += fmt.Sprintf("%.1f*%.1f,", cards[i], cards[j])
					case SUBTRACT:
						list2 = append(list2, cards[i]-cards[j])
						path += fmt.Sprintf("%.1f-%.1f,", cards[i], cards[j])
					case DIVIDE:
						if math.Abs(cards[j]) < EPSILON {
							continue
						} else {
							list2 = append(list2, cards[i]/cards[j])
							path += fmt.Sprintf("%.1f/%.1f,", cards[i], cards[j])
						}
					}
					backtracking(list2)
					list2 = list2[:len(list2)-1]
					path = oldPath
				}
			}
		}
	}
}
