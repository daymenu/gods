package insertion

// Insertion 标准插入排序
func Insertion(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

// MyInsertion 自己觉得可行的插入排序
// 经测试发现比标准快3 ns
func MyInsertion(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			} else {
				break
			}
		}
	}
}
