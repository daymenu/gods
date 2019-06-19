package strand

// Index 返回子串T在主串S中的位置，若不存在，则函数返回值为-1
func Index(S string, T string) (index int) {
	sRune, tRune := []rune(S), []rune(T)
	sLen, tLen := len(sRune), len(tRune)
	index = -1
	// 先判断字符串长度
	if sLen < tLen {
		return
	}
	maxLen := sLen - tLen
	for i := 0; i < sLen; i++ {
		if i > maxLen {
			return
		}
		for j := 0; j < tLen; j++ {
			if sRune[i+j] != tRune[j] {
				break
			}
			if j == tLen-1 {
				return i
			}
		}
	}
	return
}

func getNext(T []rune) []int {
	next := make([]int, len(T))
	next[0] = 0

	i, j := 2, 0
	for i < len(T) {
		if j == 0 || T[i-1] == T[j] {
			j++
			next[i] = j
			if T[i-1] == T[j-1] {
				next[i]++
			}
			i++
		} else {
			j = next[j-1]
		}
	}
	return next
}

// IndexKMP no finished
func IndexKMP(S string, T string) (index int) {
	sRune, tRune := []rune(S), []rune(T)
	sLen, tLen := len(sRune), len(tRune)
	index = -1
	// 先判断字符串长度
	if sLen < tLen {
		return
	}
	i, j := 0, 1
	next := getNext(tRune)
	for i <= sLen && j <= tLen {
		if j == 0 || S[i] == T[j] {
			i++
			j++
		} else {
			j = next[j-1]
		}
	}
	if j > tLen {
		return i - tLen
	}
	return
}
