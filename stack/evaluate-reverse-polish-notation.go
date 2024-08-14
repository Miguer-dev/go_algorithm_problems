package main

import (
	"strconv"
)

func operation(num1 int, num2 int, sign string) int {
	var result int

	switch sign {
	case "+":
		result = num2 + num1
	case "-":
		result = num2 - num1
	case "*":
		result = num2 * num1
	case "/":
		result = num2 / num1
	}

	return result
}

func evalRPN(tokens []string) int {
	result := 0
	stack := LIFO{}

	for _, value := range tokens {
		// is Number
		if num, err := strconv.Atoi(value); err == nil {
			if len(tokens) == 1 {
				result = num
			} else {
				stack.push(num)
			}
			// is Sign
		} else {
			result = operation(stack.pop().(int), stack.pop().(int), value)
			stack.push(result)
		}
	}

	return result
}
