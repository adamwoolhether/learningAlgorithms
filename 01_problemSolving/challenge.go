package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"sort"
	"strings"
	"unicode"
)

// 1.
func isPalindrome(w string) bool {
	i := 0
	j := len(w) - 1

	for i < j {
		if w[i] != w[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// 2.
func isPalindromeLettersOnly(w string) bool {
	i := 0
	j := len(w) - 1

	for i < j {
		for !unicode.IsLetter(rune(w[i])) {
			i++
		}
		for !unicode.IsLetter(rune(w[j])) {
			j--
		}
		if unicode.ToLower(rune(w[i])) != unicode.ToLower(rune(w[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func isPalindromeLettersOnly1(w string) bool {
	i := 0
	j := len(w) - 1

	for i < j {
		for !unicode.IsLetter(rune(w[i])) {
			i++
		}

		for !unicode.IsLetter(rune(w[j])) {
			j--
		}
		if !strings.EqualFold(string(w[i]), string(w[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func isPalindrome1(w string) bool {
	idx := 0
	for i := len(w) - 1; i > -1; i-- {
		if w[idx] != w[i] {
			return false
		}
		idx++
	}

	return true
}

func isPalindrome2(w string) bool {
	for len(w) > 1 {
		if w[0] != w[len(w)-1] {
			return false
		}
		w = w[1 : len(w)-1]
	}

	return true
}

var invCounter int

func partition(a []int, lo, hi, idx int) int {
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

func linearMedian(a []int) int {
	lo := 0
	hi := len(a) - 1
	mid := hi / 2

	for lo < hi {
		var b [8]byte
		_, _ = crand.Read(b[:])
		rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
		idx := rand.Intn(len(a))

		j := partition(a, lo, hi, idx)

		if j == mid {
			// fmt.Printf("less-than invocations: %d\n", invCounter)
			return a[j]
		}
		if j < mid {
			lo = j + 1
		} else {
			hi = j - 1
		}
	}
	// fmt.Printf("less-than invocations: %d\n", invCounter)

	return a[lo]
}

func linearMedianSorted(a []int) int {
	lo := 0
	hi := len(a) - 1
	mid := hi / 2

	sort.Ints(a)

	for lo < hi {
		var b [8]byte
		_, _ = crand.Read(b[:])
		rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
		idx := rand.Intn(len(a))

		j := partition(a, lo, hi, idx)

		if j == mid {
			// fmt.Printf("less-than invocations: %d\n", invCounter)
			return a[j]
		}
		if j < mid {
			lo = j + 1
		} else {
			hi = j - 1
		}
	}
	// fmt.Printf("less-than invocations: %d\n", invCounter)

	return a[lo]
}

// I haven't fully completed the exercises 3-5
// 3.
func countingSort(data []int, m int) []int {
	counts := make([]int, m+1)

	sortedIdx := 0

	for i := range data {
		counts[data[i]]++
	}

	for j := 0; j < m+1; j++ {
		for counts[j] > 0 {
			data[sortedIdx] = j
			sortedIdx++
			counts[j]--
		}
	}

	return data
}

