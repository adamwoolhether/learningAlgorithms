package main

import "fmt"

/*
For ascending lists, just reverse the way low and high are updated in the
for loop.
*/

func main() {
	list := []int{2, 14, 15, 19, 26, 53, 58}

	fmt.Printf("binaryArraySearch: %v\n", binaryArraySearch(list, 53))
	fmt.Printf("binaryArraySearch: %v\n", binaryArraySearch(list, 17))

	fmt.Printf("binaryArraySearchIndex: %v\n", binaryArraySearchIndex(list, 53))
	fmt.Printf("binaryArraySearchIndex: %v\n", binaryArraySearchIndex(list, 17))

	fmt.Printf("binaryArraySearchIndexOptimized: %v\n", binaryArraySearchIndexOptimized(list, 53))
	fmt.Printf("binaryArraySearchIndexOptimized: %v\n", binaryArraySearchIndexOptimized(list, 17))
}

// binaryArraySearch returns true if the given target exists within
// the given list, false if it doesn't.
func binaryArraySearch(a []int, target int) bool {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if target < a[mid] {
			high = mid - 1
		} else if target > a[mid] {
			low = mid + 1
		} else {
			return true
		}
	}

	return false
}

// binaryArraySearchIndex returns the index of given target within
// the given list. If it is not present, it returns the index of
// where the target would exist in the list, if it were to be,
// inserted, indicated with a negative number.
func binaryArraySearchIndex(a []int, target int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if target < a[mid] {
			high = mid - 1
		} else if target > a[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -(low) // book uses -(low + 1) but this seems incorrect
}

// binaryArraySearchIndexOptimized optimizes the above algorith, eliminating
// the first comparison operation with a new diff var.
//
// NOTE: This actually performs slightly worse per my benchmarks
func binaryArraySearchIndexOptimized(a []int, target int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2
		diff := target - a[mid]

		if diff < 0 {
			high = mid - 1
		} else if diff > 0 {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -(low)
}
