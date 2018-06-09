package main

import "fmt"

var heap []int

func sink(array []int, lenn int) []int {
	alen := len(array)
	k := 0
	for ; k < lenn; k++ {
		// the idea is that value in index x shall always larger than 2x+1 2x+2
		if array[k] < array[2*k+1] {
			array[k], array[2*k+1] = array[2*k+1], array[k]
		}
		if alen >= (2*k+1) && array[k] < array[2*k+2] {
			array[k], array[2*k+2] = array[2*k+2], array[k]
		}
	}
	// last circle should attent to the add and even index
	if alen == 2*k+3 {
		if array[k] < array[2*k+1] {
			array[k], array[2*k+1] = array[2*k+1], array[k]
		}
		if alen >= (2*k+1) && array[k] < array[2*k+2] {
			array[k], array[2*k+2] = array[2*k+2], array[k]
		}

	} else if array[k] < array[2*k+1] {

		array[k], array[2*k+1] = array[2*k+1], array[k]
	}

	if lenn > 1 {
		sink(array, lenn-1)
	}

	return array
}

func main() {
	orignal := []int{2, 111, 88, 901, 77, 123, 99, 2, 34, 57, 89, 11, 24, 3, 5, 11, 13, 9}
	fmt.Printf("before heap is ")
	for _, v := range orignal {
		fmt.Printf(" %v ", v)
	}
	fmt.Printf(" completed \n")
	index := len(orignal)/2 - 1

	heap = sink(orignal, index)
	fmt.Printf("after heap is ")
	for _, v := range heap {
		fmt.Printf(" %v ", v)
	}
	fmt.Printf(" completed \n")
}
