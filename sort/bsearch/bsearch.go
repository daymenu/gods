package bsearch

import "fmt"

//Bsearch 二分查找
//这个适用于不重复的递增数组
func Bsearch(a []int, value int) (int, error) {
	low := 0
	hight := len(a) - 1

	for low <= hight {
		middle := (low + hight) >> 1
		switch {
		case value == a[middle]:
			return middle, nil
		case value > a[middle]:
			low = middle + 1
		case value < a[middle]:
			hight = middle - 1
		}
	}

	return -1, fmt.Errorf("no found value")
}
