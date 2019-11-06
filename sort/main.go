package main

import (
	"fmt"

	"github.com/daymenu/gods/sort/bubble"
)

func main() {
	a := []int{3, 2, 1, 5, 9, 7}
	bubble.Bubble(a)
	fmt.Println(a)
}
