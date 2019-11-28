package bsearch

import "fmt"

// RepeatLBsearchRecurse 有重复元素的二分查找最左边的元素
func RepeatLBsearchRecurse(a []int, value int) (i int, err error) {
	return rlBsearchRecurse(a, value, 0, len(a)-1)
}

func rlBsearchRecurse(a []int, value, low, hight int) (i int, err error) {
	if low > hight {
		return -1, fmt.Errorf("Bsearch %d not in a", value)
	}
	mid := low + ((hight - low) >> 1)
	switch {
	case a[mid] > value:
		return rlBsearchRecurse(a, value, low, mid-1)
	case a[mid] < value:
		return rlBsearchRecurse(a, value, mid+1, hight)
	case a[mid] == value:
		if mid > 0 && a[mid-1] == value {
			return rlBsearchRecurse(a, value, low, mid-1)
		}
		return mid, nil
	}
	return -1, fmt.Errorf("没有找到元素哦")
}

// RepeatRBsearchRecurse 有重复元素的二分查找最右边的元素
func RepeatRBsearchRecurse(a []int, value int) (i int, err error) {
	return rrBsearchRecurse(a, value, 0, len(a)-1)
}

func rrBsearchRecurse(a []int, value, low, hight int) (i int, err error) {
	if low > hight {
		return -1, fmt.Errorf("Bsearch %d not in a", value)
	}
	mid := low + ((hight - low) >> 1)
	switch {
	case a[mid] > value:
		return rrBsearchRecurse(a, value, low, mid-1)
	case a[mid] < value:
		return rrBsearchRecurse(a, value, mid+1, hight)
	case a[mid] == value:
		if mid+1 < len(a) && a[mid+1] == value {
			return rrBsearchRecurse(a, value, mid+1, hight)
		}
		return mid, nil
	}
	return -1, fmt.Errorf("没有找到元素哦")
}
