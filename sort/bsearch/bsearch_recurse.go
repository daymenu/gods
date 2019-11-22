package bsearch

import "fmt"

//Recurse 二分查找
func Recurse(a []int, value int) (int, error) {
	return bsearchRecurse(a, value, 0, len(a)-1)
}

func bsearchRecurse(a []int, value, min, max int) (int, error) {
	if min > max {
		return -1, fmt.Errorf("Bsearch %d not in a", value)
	}
	middle := (max + min) / 2
	if a[middle] == value {
		return middle, nil
	}
	if value > a[middle] {
		return bsearchRecurse(a, value, middle+1, max)
	}
	return bsearchRecurse(a, value, min, middle-1)

}
