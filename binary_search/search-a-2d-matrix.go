package binary_search

func SearchMatrix(matrix [][]int, target int) bool {

	for _, array := range matrix {
		if search(array, target) {
			return true
		}
	}

	return false
}

func search(nums []int, target int) bool {
	if len(nums) != 0 {
		if nums[len(nums)/2] == target {
			return true
		} else if nums[len(nums)/2] > target {
			return search(nums[:len(nums)/2], target)
		} else if nums[len(nums)/2] < target {
			return search(nums[len(nums)/2+1:], target)
		}
	}

	return false
}
