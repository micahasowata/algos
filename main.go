package main

import (
	"algos/bs"
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 2, 4, 5, 7, 3, 2, 9}
	slices.Sort(nums)           // sort the array
	nums = slices.Compact(nums) // remove duplicate elements

	fmt.Println(nums)
	target := 0

	result := bs.Search(nums, target)
	if result == -1 {
		panic("target does not exist in array")
	}

	fmt.Printf("%d is at position %d", target, result)
}
