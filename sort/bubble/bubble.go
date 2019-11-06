package bubble

// Bubble 冒泡排序算法
func Bubble(a []int) []int {
	for i := range a {
		swap := false
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
			swap = true
		}
		if !swap {
			break
		}
	}
	return a
}
