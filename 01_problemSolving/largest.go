package main

import (
	"errors"
)

// Our largest func differs slight from the book.
// We also return the max number's index, because removing a value from a slice
// in go requires knowing the index position. We use this later in doubleTwo
func largest(a ...int) (int, int, error) {
	if len(a) == 0 {
		return 0, 0, errors.New("slice must not be empty")
	}

	max, maxIndex := a[0], 0
	for i, v := range a[1:] {
		if max < v {
			max, maxIndex = v, i+1
		}
	}

	return max, maxIndex, nil
}

func alternate(a ...int) int {
	for _, v := range a {
		vIsMax := true

		for _, v2 := range a {
			if v < v2 {
				vIsMax = false
				break
			}
		}
		if vIsMax {
			return v
		}
	}

	return 0
}
