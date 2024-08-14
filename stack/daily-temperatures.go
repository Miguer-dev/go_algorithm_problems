package main

func dailyTemperatures(temps []int) []int {
	results := make([]int, len(temps))
	stack := LIFO{}

	for index, value := range temps {

		for len(stack) > 0 && temps[stack[len(stack)-1].(int)] < value {
			lastStackValue := stack.pop().(int)
			results[lastStackValue] = index - lastStackValue
		}
		stack.push(index)
	}

	return results
}
