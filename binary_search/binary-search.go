package binary_search

func BinarySearch(nums []int, target int, passlen int) int {
	if len(nums) != 0 {
		if nums[len(nums)/2] == target {
			return len(nums)/2 + passlen
		} else if nums[len(nums)/2] > target {
			return BinarySearch(nums[:len(nums)/2], target, passlen)
		} else if nums[len(nums)/2] < target {
			return BinarySearch(nums[len(nums)/2+1:], target, (len(nums)/2+1)+passlen)
		}
	}

	return -1
}
