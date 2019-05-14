package list

import (
	"testing"
)

func TestGetElem(t *testing.T) {
	initList()
	s, err := GetElem(1)
	if err != nil {
		t.Error(err)
	}

	if s.No != students[0].No {
		t.Errorf("%s = %s, want %s", s.No, students[0].No, s.No)
	}
}

func TestListInsert(t *testing.T) {
	clearList()
	ListInsert(1, Student{"201211032012", "gsy", "男"})
	if len(students) != 1 {
		t.Errorf("list length is %d, want %d", len(students), 1)
	}
	ListInsert(2, Student{"201211032012", "hj", "女"})
	if len(students) != 2 {
		t.Errorf("list length is %d, want %d", len(students), 2)
	}
}

func TestListDelete(t *testing.T) {
	initList()
	s, err := ListDelete(1)
	if err != nil {
		t.Error(err)
	}
	want := "201211032011"
	if s.No != "201211032011" {
		t.Errorf("No = %s, want %s", s.No, want)
	}
	if len(students) != 1 {
		t.Errorf("list length is %d, want %d", len(students), 1)
	}
}

func TestLocateElem(t *testing.T) {
	initList()
	no := "201211032011"
	i, ok := LocateElem(no)
	if !ok {
		t.Errorf("can't find student:%s", no)
	}

	if i != 1 {
		t.Errorf("index is %d, want %d",i,1)
	}

}

func initList() {
	students = []Student{
		{"201211032011", "gsy", "男"},
		{"201211032012", "hj", "女"},
	}
}

func clearList() {
	students = students[:0]
}
