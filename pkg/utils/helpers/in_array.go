package helpers

import "sort"

// 检查字符串是否在数组中,有坑会改变数组的顺序

func InArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}
