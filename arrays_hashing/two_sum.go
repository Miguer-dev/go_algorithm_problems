package array_shashing

func TwoSum(nums []int, target int) []int {
	//key=value, value=array index
	numsMap := make(map[int][]int)

	for i, v := range nums {
		if _, exits := numsMap[v]; exits {
			numsMap[v] = append(numsMap[v], i)
		} else {
			numsMap[v] = []int{i}
		}
	}

	for k := range numsMap {
		findNum := target - k

		if _, exits := numsMap[findNum]; exits {

			switch {
			case findNum == k && len(numsMap[findNum]) >= 2:
				return []int{numsMap[findNum][0], numsMap[findNum][1]}
			case findNum == k && len(numsMap[findNum]) < 2:
				break
			default:
				return []int{numsMap[k][0], numsMap[findNum][0]}
			}
		}
	}

	return []int{0, 0}
}
