package main

import (
	"sort"
)

func sortStringBytes(s string) string {

	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })

	return string(bytes)
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	//map[sortString] [realString,realString,realString]
	strsSorted := map[string][]string{}

	for _, v := range strs {
		str := sortStringBytes(v)
		if _, exits := strsSorted[str]; exits {
			strsSorted[str] = append(strsSorted[str], v)
		} else {
			strsSorted[str] = []string{v}
		}
	}

	for _, v := range strsSorted {
		result = append(result, v)
	}

	return result
}

func groupAnagrams2(strs []string) [][]string {
	result := [][]string{}
	//map[sortString] resultIndex
	strMap := map[string]int{}

	for _, str := range strs {
		strSort := sortStringBytes(str)
		if i, exits := strMap[strSort]; exits {
			result[i] = append(result[i], str)
		} else {
			result = append(result, []string{str})
			strMap[strSort] = len(result) - 1
		}
	}

	return result
}
