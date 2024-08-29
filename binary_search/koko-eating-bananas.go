package binary_search

import (
	"math"
	"slices"
)

func MinEatingSpeed(piles []int, h int) int {
	if len(piles) == 1 {
		return int(math.Ceil(float64(piles[0]) / float64(h)))
	}

	minPile := slices.Min(piles)
	maxPile := slices.Max(piles)

	return binarySearchPiles(minPile, maxPile, piles, h)

}

func binarySearchPiles(min int, max int, piles []int, hours int) int {
	if min == max {
		if check(piles, hours, min) {
			return min
		} else {
			return min + 1
		}

	}

	mediumSpeed := (min + max) / 2

	if check(piles, hours, mediumSpeed) {
		return binarySearchPiles(min, mediumSpeed-1, piles, hours)
	} else if !check(piles, hours, mediumSpeed) {
		return binarySearchPiles(mediumSpeed+1, max, piles, hours)
	}

	return -1
}

func check(piles []int, hours int, speed int) bool {
	result := 0

	for _, pile := range piles {
		result += int(math.Ceil(float64(pile) / float64(speed)))

		if result > hours {
			return false
		}
	}

	return true
}
