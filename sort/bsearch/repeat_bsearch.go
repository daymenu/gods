package bsearch

import "fmt"

// RepeatLBsearch 有重复元素的二分查找最左边的元素
func RepeatLBsearch(a []int, value int) (i int, err error) {
	low, hight := 0, len(a)-1

	for low <= hight {
		mid := (low + hight) >> 1
		fmt.Printf("low=%d hight=%d mid=%d a[%d]=%d value=%d \n", low, hight, mid, mid, a[mid], value)
		switch {
		case a[mid] < value:
			low = mid + 1
		case a[mid] == value:
			if mid-1 >= 0 && a[mid-1] == value {
				hight = mid - 1
				continue
			}
			return mid, nil
		case a[mid] > value:
			hight = mid - 1
		}
	}
	return -1, fmt.Errorf("no found value")
}

// RepeatRBsearch 有重复元素的二分查找最右边的元素
func RepeatRBsearch(a []int, value int) (i int, err error) {
	low, hight := 0, len(a)-1

	for low <= hight {
		mid := (low + hight) >> 1
		switch {
		case a[mid] < value:
			low = mid + 1
		case a[mid] == value:
			if mid+1 < len(a) && a[mid+1] == value {
				low = mid + 1
				continue
			}
			return mid, nil
		case a[mid] > value:
			hight = mid - 1
		}
	}
	return -1, fmt.Errorf("no found value")
}
