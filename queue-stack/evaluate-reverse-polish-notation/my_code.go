package main

import (
	"fmt"
	"strconv"
)

func main() {
	token := []string{"2", "1", "+", "3", "*"}

	fmt.Println(evalRPN(token))
}

func evalRPN(tokens []string) int {
	var test []int

	for i := range tokens {
		if num, err := strconv.Atoi(tokens[i]); err != nil {
			num1 := test[len(test)-2]
			num2 := test[len(test)-1]
			var tmp int
			switch tokens[i] {
			case "+":
				tmp = num1 + num2
				test = test[:len(test)-2]
				test = append(test, tmp)
			case "-":
				tmp = num1 - num2
				test = test[:len(test)-2]
				test = append(test, tmp)
			case "*":
				tmp = num1 * num2
				test = test[:len(test)-2]
				test = append(test, tmp)
			case "/":
				tmp = num1 / num2
				test = test[:len(test)-2]
				test = append(test, tmp)
			}
		} else {
			test = append(test, num)
		}
	}

	return test[0]
}
