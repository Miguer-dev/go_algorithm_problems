package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	var backValue int
	count := 0
	maxCount := 0
	for index, value := range nums {

		if value-1 == backValue || index == 0 {
			count++
		} else if value == backValue {
			continue
		} else {
			count = 1
		}

		if count > maxCount {
			maxCount = count
		}

		backValue = value
	}

	return maxCount
}
