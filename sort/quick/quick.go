package quick

// Quick  快排
func Quick(a []int) {
	split(a, 0, len(a)-1)
}

func partition(a []int, p, r int) int {
	i := p
	for ; p < r; p++ {
		if a[p] < a[r] {
			a[p], a[i] = a[i], a[p]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}

func split(a []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(a, p, r)
	split(a, p, q-1)
	split(a, q+1, r)
}
