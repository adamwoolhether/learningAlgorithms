package main

import "fmt"

/*
Note to self: remember to compare this to the commented out implementation below (curious)
*/

func insertionSort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if a[j-1] <= a[j] {
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func main() {
	array := []int{4, 76, 3, 223, 9, 0, 69, 98, 3, 62}
	fmt.Println(array)

	insertionSort(array)
	fmt.Println(array)
}

/*
func InsertionSortInt(list []int) {
	var sortedList []int

	for _, listValue := range list {
		sortedList = func(sortedList []int, listValue int) []int {
			for i, sortValue := range sortedList {
				if listValue < sortValue {
					// sortedList[:i] + listValue + sotedList[i:]
					return append(sortedList[:i], append([]int{listValue}, sortedList[i:]...)...)
				}
			}
			return append(sortedList, listValue)
		}(sortedList, listValue)
	}

	for i, v := range sortedList {
		list[i] = v
	}
}
*/
