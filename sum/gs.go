package sum

// 高斯求和
func GsSum(start int, end int) (s int) {
	s = (start + end) * (end - start + 1) / 2
	return s
}
