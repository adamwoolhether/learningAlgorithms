package main

import (
	"fmt"
	"os"
	"sort"
)

func largestTwo(a ...int) (int, int) {
	if len(a) < 2 {
		fmt.Println("must have > 2 values")
		os.Exit(1)
	}

	max, second := a[0], a[1]
	if max < second {
		max, second = second, max
	}

	for _, v := range a[2:] {
		if max < v {
			max, second = v, max
		} else if second < v {
			second = v
		}
	}

	return max, second
}

func sortingTwo(a ...int) (int, int) {
	if len(a) < 2 {
		fmt.Println("must have > 2 values")
		os.Exit(1)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	return a[0], a[1]
}

func doubleTwo(a ...int) (int, int) {
	if len(a) < 2 {
		fmt.Println("must have > 2 values")
		os.Exit(1)
	}

	max, maxIndex, _ := largest(a...)

	aCopy := make([]int, len(a))
	copy(aCopy, a)

	aCopy = append(aCopy[:maxIndex], aCopy[maxIndex+1:]...)

	second, _, _ := largest(aCopy...)

	return max, second
}

func mutableTwo(a ...int) (int, int) {
	if len(a) < 2 {
		fmt.Println("must have > 2 values")
		os.Exit(1)
	}

	max, maxIndex, _ := largest(a...)

	a = append(a[:maxIndex], a[maxIndex+1:]...)

	second, _, _ := largest(a...)

	a = append(a[:maxIndex], append([]int{max}, a[maxIndex:]...)...)

	return max, second
}

func tournamentTwo(a ...int) (int, int) {
	n := len(a)
	winner := make([]int, n-1)
	loser := make([]int, n-1)
	prior := make([]int, n-1)
	prior[0] = -1

	idx := 0
	for i := 0; i < n; i += 2 {
		if a[i] < a[i+1] {
			winner[idx] = a[i+1]
			loser[idx] = a[i]
		} else {
			winner[idx] = a[i]
			loser[idx] = a[i+1]
		}
		idx++
	}

	m := 0
	for idx < n-1 {
		if winner[m] < winner[m+1] {
			winner[idx] = winner[m+1]
			loser[idx] = winner[m]
			prior[idx] = m + 1
		} else {
			winner[idx] = winner[m]
			loser[idx] = winner[m+1]
			prior[idx] = m
		}
		m += 2
		idx++
	}

	largest := winner[m]
	second := loser[m]

	for m = prior[m]; m >= 0; {
		if second < loser[m] {
			second = loser[m]
		}
		m = prior[m]
	}

	return largest, second
}
