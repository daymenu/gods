package list

type LStudent struct {
	stu  Student
	next *LStudent
}

type List struct {
	head   *LStudent
	length int
}

func InitLinkList() *List {
	return &List{
		head: &LStudent{next: nil},
	}
}

func (l *List) GetLinkElem(i int) (s *Student, err error) {
	p := l.head.next
	for j := 1; j <= i; j++ {
		if p == nil {
			break
		}
		if j == i {
			s = &p.stu
			break
		}
		p = p.next
	}
	return
}

func (l *List) ListLinkLength() int {
	return l.length
}

func (l *List) ListLinkInsert(i int, stu Student) error {
	if i > l.ListLinkLength()+1 || i < 1 {
		return IndexError
	}
	pstu := LStudent{stu: stu}
	p := l.head
	for j := 1; j <= i; j++ {
		if p == nil {
			break
		}
		if j == i {
			pstu.next = p.next
			p.next = &pstu
			l.length++
			break
		}
		p = p.next
	}
	return nil
}

func (l *List) ListLinkDelete(i int) (stu *Student, err error) {
	if i > l.ListLinkLength() || i < 1 {
		return nil, IndexError
	}
	p := l.head
	for j := 1; j <= i; j++ {
		if p == nil {
			break
		}
		if j == i {
			stu = &p.stu
			p.next = p.next.next
			l.length--
			break
		}
		p = p.next
	}
	return
}
