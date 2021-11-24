package main

import "strings"

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


