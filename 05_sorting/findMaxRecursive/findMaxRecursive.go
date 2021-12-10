package main

import "fmt"

func findMax(a []int) int {
	var rmax func(low, high int) int

	rmax = func(low, high int) int {
		if low == high {
			return a[low]
		}
		mid := (low + high) / 2
		l := rmax(low, mid)
		r := rmax(mid+1, high)

		return max(l, r)
	}

	return rmax(0, len(a)-1)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {
	array := []int{4, 76, 3, 223, 90, 69, 98, 3, 62}
	fmt.Println(findMax(array))
}
