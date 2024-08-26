package two_pointers

func twoSum(numbers []int, target int) []int {
	result := []int{}

	pointerStart := 0
	pointerEnd := len(numbers) - 1

	for pointerStart < pointerEnd {

		sum := numbers[pointerStart] + numbers[pointerEnd]

		if sum == target {
			result = append(result, pointerStart+1, pointerEnd+1)
			break
		}
		if sum < target {
			pointerStart += 1
			continue
		}
		if sum > target {
			pointerEnd -= 1
			continue
		}

	}

	return result
}
