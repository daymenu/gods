package merge

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
	return merge(split(a[:aLen/2]), split(a[aLen/2:]))
}

func merge(a, b []int) []int {
	aLen, bLen := len(a), len(b)
	c := make([]int, aLen+bLen)
	i, j, k := 0, 0, 0
	for i < aLen && j < bLen {
		if a[i] <= b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
		k++
	}
	if i < aLen {
		for ; i < aLen; i++ {
			c[k] = a[i]
			k++
		}
	}

	if j < bLen {
		for ; j < bLen; j++ {
			c[k] = b[j]
			k++
		}
	}
	return c
}
