package main

import "fmt"

// Merge 归并排序
func Merge(a []int) []int {
	return split(a)
}

func split(a []int) []int {

	aLen := len(a)
	if aLen == 1 {
		return a
	}
	if aLen == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
	}
	fmt.Println("split:", a[:aLen/2], a[aLen/2:])
	return merge(split(a[:aLen/2]), split(a[aLen/2:]))
}

func merge(a, b []int) []int {
	// fmt.Println("merge", a, b)
	aLen, bLen := len(a), len(b)
	c := make([]int, aLen+bLen)
	c[0] = 1
	i, j := 0, 0
	for i < aLen && j < bLen {
		if a[i] > a[j] {
			c = append(c, a[j])
			j++
		} else {
			c = append(c, a[i])
			i++
		}
	}

	if i < aLen-1 {
		for ; i < aLen; i++ {
			c = append(c, a[i])
		}
	}

	if j < aLen-1 {
		for ; j < aLen; j++ {
			c = append(c, a[j])
		}
	}

	return c
}

func main() {
	fmt.Println(Merge([]int{3, 2, 7, 6, 8, 1, 0}))
}
