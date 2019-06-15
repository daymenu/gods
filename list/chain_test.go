package list

import (
	"testing"
)

func TestListLinkInsert(t *testing.T) {
	linklist := InitLinkList()
	linklist.ListLinkInsert(1, Student{"201211032011", "gsy", "男"})
	if linklist.head.next.stu.No != "201211032011" {
		t.Errorf("insert into student no is %s , want %s", linklist.head.next.stu.No, "201211032011")
	}
	linklist.ListLinkInsert(2, Student{"201211032012", "hj", "女"})
	if linklist.head.next.next.stu.No != "201211032012" {
		t.Errorf("insert into student no is %s , want %s", linklist.head.next.stu.No, "201211032012")
	}
}

func TestListLinkDelete(t *testing.T) {
	linklist := InitLinkList()
	linklist.ListLinkInsert(1, Student{"201211032011", "gsy", "男"})
	linklist.ListLinkInsert(2, Student{"201211032012", "hj", "女"})
	linklist.ListLinkDelete(1)
	linklist.ListLinkDelete(1)
	t.Error(linklist.head.next)
}
