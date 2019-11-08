package main

import (
	"fmt"

	"github.com/daymenu/gods/sort/bubble"
	"github.com/daymenu/gods/sort/insertion"
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
}
