package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2, 10, 8, 4}
	b := []int{3, 0, 2, 7, 4}

	sort.Ints(a)
	sort.Ints(b)

	c := append(a, b...)

	fmt.Println(c)
	fmt.Println(len(c))
}
