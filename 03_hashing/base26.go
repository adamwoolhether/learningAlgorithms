package main

import (
	"fmt"
	"strings"
)

// base26 will convert a string into its
// integer ASCII representation.
func base26(w string) int {
	val := 0
	w = strings.ToLower(w)
	for char := range w {
		nextDigit := w[char] - 'a'
		val = 26*val + int(nextDigit)
	}

	return val
}

func main() {
	January := base26("January")

	Canada := base26("Canada")

	Mike := base26("Mike")

	fmt.Println(January)
	fmt.Println(Canada)
	fmt.Println(Mike)
}
