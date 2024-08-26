package two_pointers

func twoSum(numbers []int, target int) []int {
	result := []int{}

	pointerStart := 0
	pointerEnd := len(numbers) - 1

	for pointerStart < pointerEnd {

		sum := numbers[pointerStart] + numbers[pointerEnd]

		if sum == target {
			return append(result, pointerStart+1, pointerEnd+1)
		}
		if sum < target {
			pointerStart++
		} else if sum > target {
			pointerEnd--
		}

	}

	return result
}
