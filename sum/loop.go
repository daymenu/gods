package sum

func LoopSum(start int, end int) (s int) {
	for i := start; i <= end; i++ {
		s += i
	}
	return s
}
