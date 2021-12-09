package main

import "fmt"

func selectionSort(a []int) {
	n := len(a)
	for i := range a {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}

		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

func main() {
	array := []int{4, 76, 3, 223, 9, 0, 69, 98, 3, 62}
	fmt.Println(array)

	selectionSort(array)
	fmt.Println(array)
}
