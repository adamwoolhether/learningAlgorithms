package main

import (
	"fmt"
)

func quicksort(a []int) {

	var qsort func(int, int)
	qsort = func(low, high int) {
		if high <= low {
			return
		}

		// rand.Seed(time.Now().UnixNano())
		// pivotIdx := rand.Intn(len(a))
		pivotIdx := low
		location := partition(a, low, high, pivotIdx)

		qsort(low, location-1)
		qsort(location+1, high)
	}

	qsort(0, len(a)-1)
}

func partition(a []int, lo, hi, idx int) int {
	var invCounter int

	if lo == hi {
		return lo
	}

	a[idx], a[lo] = a[lo], a[idx]
	i := lo
	j := hi + 1

	for {
		for {
			i++
			if i == hi {
				break
			}
			if a[lo] < a[i] {
				invCounter++
				break
			}
		}

		for {
			j--
			if j == lo {
				break
			}
			if a[j] < a[lo] {
				invCounter++
				break
			}
		}

		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}

	a[lo], a[j] = a[j], a[lo]

	return j
}

func main() {
	array := []int{4, 76, 3, 223, 90, 69, 98, 75, 62}
	fmt.Println(array)

	quicksort(array)
	fmt.Println(array)
}
