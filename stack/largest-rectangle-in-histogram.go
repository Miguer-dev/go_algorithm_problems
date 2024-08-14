package main

import (
	"slices"
	"sort"
)

type Height struct {
	value     int
	positions []int
}

func LargestRectangleArea(heights []int) int {
	// key = value, value = [indexes]
	heightsMap := make(map[int][]int)
	heightsSort := []Height{}

	for index, value := range heights {
		heightsMap[value] = append(heightsMap[value], index)
	}

	for key, value := range heightsMap {
		heightsSort = append(heightsSort, Height{key, value})
	}

	sort.Slice(heightsSort, func(i, j int) bool {
		return heightsSort[i].value > heightsSort[j].value
	})

	maxHeight := 0
	backHeights := []int{}
	for _, height := range heightsSort {
		height.positions = append(height.positions, backHeights...)
		slices.Sort(height.positions)

		continuousHeight := 0
		for index, value := range height.positions {
			if index == 0 || height.positions[index-1] == value-1 {
				continuousHeight++
			} else {
				continuousHeight = 1
			}

			tmpHeight := continuousHeight * height.value
			if tmpHeight > maxHeight {
				maxHeight = tmpHeight
			}
		}

		backHeights = height.positions
	}

	return maxHeight
}

func largestRectangleArea2(heights []int) int {
	totalHeight := 0

	for {
		positionsHeight := []int{}
		maxHeight := 0
		beforeMaxHeight := 0

		for index, value := range heights {

			if value > maxHeight {
				beforeMaxHeight = maxHeight
				maxHeight = value
				positionsHeight = []int{index}
			} else if value == maxHeight {
				positionsHeight = append(positionsHeight, index)
			}

			if value < maxHeight && value > beforeMaxHeight {
				beforeMaxHeight = value
			}
		}

		slices.Sort(positionsHeight)
		continuousHeight := 0
		for index, value := range positionsHeight {
			if index == 0 || positionsHeight[index-1] == value-1 {
				continuousHeight++
			} else {
				continuousHeight = 1
			}

			tmpHeight := continuousHeight * maxHeight
			if tmpHeight > totalHeight {
				totalHeight = tmpHeight
			}
		}

		for _, value := range positionsHeight {
			heights[value] = beforeMaxHeight
		}

		if len(positionsHeight) == len(heights) {
			break
		}
	}

	return totalHeight
}
