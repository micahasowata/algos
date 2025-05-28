/*
bs implements binary search
*/
package bs

// Search performs a binary search to find target in sorted array and returns its index position if // it can find it or -1 if it can't
//
// x must be sorted and have only unique items for this algorithm to perform as intended
func Search(sortedArr []int, target int) int {
	lowest, highest := 0, len(sortedArr)-1

	for lowest < highest {
		center := (lowest + highest) / 2
		if target == sortedArr[center] {
			return center
		}

		// target is lower than the center of the sorted array
		if target < sortedArr[center] {
			highest = center - 1
		}

		// target is higher than the center of the sorted array
		if target > sortedArr[center] {
			lowest = center + 1
		}
	}

	return -1
}
