package main

import (
	"fmt"

	"github.com/daymenu/gods/sort/bubble"
	"github.com/daymenu/gods/sort/insertion"
	"github.com/daymenu/gods/sort/merge"
	"github.com/daymenu/gods/sort/quick"
	"github.com/daymenu/gods/sort/selection"
)

func main() {
	a := []int{3, 2, 1, 5, 9, 7}
	bubble.Bubble(a)
	fmt.Println(a)
	a = []int{3, 2, 1, 5, 9, 7}
	insertion.Insertion(a)
	fmt.Println(a)
	a = []int{3, 2, 1, 5, 9, 7}
	selection.Selection(a)
	fmt.Println(a)
	a = []int{3, 2, 1, 1, 5, 6, 10, 9, 7}
	quick.Quick(a)
	fmt.Println(a)
	a = []int{3, 2, 1, 1, 5, 6, 10, 9, 7}
	b := merge.Merge(a)
	fmt.Println(b)
}
