package bsearch

import (
  "testing"
)

func TestBsearchResurse(t *testing.T){
  a = []int{1,2,3,4,5,6,7,8}
  i,e:=BsearchRecurse(a)
  if e!= nil{
    t.Errorf(a)
  }
}
