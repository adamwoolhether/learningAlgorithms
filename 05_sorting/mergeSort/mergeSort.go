package main

import "fmt"

func mergeSort(a []int) {
	aux := make([]int, len(a))

	merge := func(low, mid, high int) {
		copy(aux[low:high+1], a[low:high+1])

		left, right := low, mid+1

		for i := low; i < high; i++ {
			switch {
			case left > mid:
				a[i] = aux[right]
				right++
			case right > high:
				a[i] = aux[left]
				left++
			case aux[right] < aux[left]:
				a[i] = aux[right]
				right++
			default:
				a[i] = aux[left]
				left++
			}
		}
	}

	var rsort func(int, int)
	rsort = func(low, high int) {
		if high <= low {
			return
		}
		mid := (low + high) / 2

		rsort(low, mid)
		rsort(mid+1, high)
		merge(low, mid, high)
	}

	rsort(0, len(a)-1)
}

func main() {
	array := []int{4, 76, 3, 223, 90, 69, 98, 75, 62}
	fmt.Println(array)

	mergeSort(array)
	fmt.Println(array)
}
