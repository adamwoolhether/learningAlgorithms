package main

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		input string
		is    bool
	}{
		{"madam", true},
		{"racecar", true},
		{"random", false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("string: %s", tc.input), func(*testing.T) {
			got := isPalindrome(tc.input)
			if got != tc.is {
				t.Fatalf("isPalindrome() = %v; want %v", got, tc.is)
			}
		})
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("string: %s", tc.input), func(*testing.T) {
			got := isPalindrome1(tc.input)
			if got != tc.is {
				t.Fatalf("isPalindrome1() = %v; want %v", got, tc.is)
			}
		})
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("string: %s", tc.input), func(*testing.T) {
			got := isPalindrome2(tc.input)
			if got != tc.is {
				t.Fatalf("isPalindrome1() = %v; want %v", got, tc.is)
			}
		})
	}
}

func Benchmark_isPalindrome_redivider(t *testing.B) {
	w := "redivider"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindrome(w)
	}
}
func Benchmark_isPalindrome_tattarrattat(t *testing.B) {
	w := "tattarrattat"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindrome(w)
	}
}
func Benchmark_isPalindrome_nomelgibsonisacasinosbiglemon(t *testing.B) {
	w := "nomelgibsonisacasinosbiglemon"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindrome(w)
	}
}

// =====================================================================================================================
func Benchmark_isPalindrome1_redivider(t *testing.B) {
	w := "redivider"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindrome1(w)
	}
}
func Benchmark_isPalindrome1_tattarrattat(t *testing.B) {
	w := "tattarrattat"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindrome1(w)
	}
}
func Benchmark_isPalindrome1_nomelgibsonisacasinosbiglemon(t *testing.B) {
	w := "nomelgibsonisacasinosbiglemon"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindrome1(w)
	}
}

// =====================================================================================================================
func Benchmark_isPalindrome2_redivider(t *testing.B) {
	w := "redivider"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindrome2(w)
	}
}
func Benchmark_isPalindrome2_tattarrattat(t *testing.B) {
	w := "tattarrattat"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindrome2(w)
	}
}
func Benchmark_isPalindrome2_nomelgibsonisacasinosbiglemon(t *testing.B) {
	w := "nomelgibsonisacasinosbiglemon"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindrome2(w)
	}
}

// =====================================================================================================================
func Test_isPalindromeLettersOnly(t *testing.T) {
	tests2 := []struct {
		input string
		is    bool
	}{
		{"Able was I ere I saw Elba", true},
		{"A man, a plan, a canal - Panama", true},
		{"Madam, I'm Adam", true},
		{"Never odd or even", true},
		{"Doc, note: I dissent. A fast never prevents a fatness. I diet on cod", true},
		{"T. Eliot, top bard, notes putrid tang emanating, is sad; I'd assign it a name: gnat dirt upset on drab pot toilet.", true},
		{"The quick brown fox jumps over the lazy dog", false},
	}

	for _, tc2 := range tests2 {
		t.Run(fmt.Sprintf("string: %s", tc2.input), func(*testing.T) {
			got := isPalindromeLettersOnly(tc2.input)
			if got != tc2.is {
				t.Fatalf("isPalindromLettersOnly() = %v; want %v", got, tc2.is)
			}
		})
	}

	for _, tc2 := range tests2 {
		t.Run(fmt.Sprintf("string: %s", tc2.input), func(*testing.T) {
			got := isPalindromeLettersOnly1(tc2.input)
			if got != tc2.is {
				t.Fatalf("isPalindromLettersOnly1() = %v; want %v", got, tc2.is)
			}
		})
	}
}

func Benchmark_isPalindromeLettersOnly_short(t *testing.B) {
	w := "Madam, I'm Adam"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindromeLettersOnly(w)
	}

}
func Benchmark_isPalindromeLettersOnly_long(t *testing.B) {
	w := "Doc, note: I dissent. A fast never prevents a fatness. I diet on cod"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindromeLettersOnly(w)
	}
}

// =====================================================================================================================
func Benchmark_isPalindromeLettersOnly1_short(t *testing.B) {
	w := "Madam, I'm Adam"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindromeLettersOnly1(w)
	}

}
func Benchmark_isPalindromeLettersOnly1_long(t *testing.B) {
	w := "Doc, note: I dissent. A fast never prevents a fatness. I diet on cod"
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		isPalindromeLettersOnly1(w)
	}
}

// =====================================================================================================================
func Benchmark_linearMedian_10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedian(slice)
	}
}
func Benchmark_linearMedian_100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedian(slice)
	}
}
func Benchmark_linearMedian_1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedian(slice)
	}
}
func Benchmark_linearMedian_10K(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedian(slice)
	}
}
func Benchmark_linearMedian_100K(t *testing.B) {
	slice := ints[0:100000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedian(slice)
	}
}
func Benchmark_linearMedian_1000K(t *testing.B) {
	slice := ints[0:1000000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedian(slice)
	}
}


func Benchmark_linearMedianSorted_10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedianSorted(slice)
	}
}
func Benchmark_linearMedianSorted_100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedianSorted(slice)
	}
}
func Benchmark_linearMedianSorted_1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedianSorted(slice)
	}
}
func Benchmark_linearMedianSorted_10K(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedianSorted(slice)
	}
}
func Benchmark_linearMedianSorted_100K(t *testing.B) {
	slice := ints[0:100000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedianSorted(slice)
	}
}
func Benchmark_linearMedianSorted_1000K(t *testing.B) {
	slice := ints[0:1000000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		linearMedianSorted(slice)
	}
}

// =====================================================================================================================
