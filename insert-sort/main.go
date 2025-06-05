package main

import "fmt"

// insertSort performs an insertion sort on A
func insertSort(A []int) []int {
	for i := 1; i < len(A); i++ {
		current := A[i]
		prev := i - 1

		for prev >= 0 && A[prev] > current {
			A[prev+1] = A[prev]
			prev = prev - 1
		}

		A[prev+1] = current
	}

	return A
}

func main() {
	a := insertSort([]int{5, 2, 4, 6, 1, 3})

	fmt.Println(a)
}
