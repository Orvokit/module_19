package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 6

func create() (a []int) {

	a = make([]int, size)
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		a[i] = rand.Intn(100)
	}
	return
}

func bubbleSort(a []int) []int {
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func main() {
	a := create()
	fmt.Println("Сгенерированный начальный массив:", a)
	sortedA := bubbleSort(a)
	fmt.Println("Отсортированный пузырьком массив:", sortedA)
}
