package main

func productExceptSelf(nums []int) []int {
	result := []int{}
	total := 1
	cantZeros := 0

	for _, value := range nums {
		if value == 0 {
			cantZeros++
			if cantZeros > 1 {
				break
			}
		} else {
			total *= value
		}
	}

	for _, value := range nums {
		switch {
		case cantZeros > 1:
			result = append(result, 0)
		case value == 0:
			result = append(result, total)
		case value != 0 && cantZeros == 1:
			result = append(result, 0)
		default:
			result = append(result, total/value)
		}
	}

	return result
}
