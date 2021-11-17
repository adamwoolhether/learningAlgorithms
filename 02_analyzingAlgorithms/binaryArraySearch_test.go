package main

import (
	"math/rand"
	"fmt"
	"testing"
	"time"
)

var ints []int

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	ints = make([]int, 1_000_000)

	for i := 0; i < 1_000_000; i++ {
		ints[i] = rand.Int()
	}

	m.Run()
}

// =====================================================================================================================
// binaryArraySearch
// =====================================================================================================================

func Test_binaryArraySearch(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		result bool
	}{
		{[]int{2, 14, 15, 19, 26, 53, 58}, 53, true},
		{[]int{2, 14, 15, 19, 26, 53, 58}, 17, false},
		{[]int{-99, -60, -4, -1}, -60, true},
		{[]int{-99, -60, -4, -1}, 60, false},
		{[]int{-99, -60, -4, -1, 0, 30, 40, 99, 4999}, 30, true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got := binaryArraySearch(tc.input, tc.target)
			if got != tc.result {
				t.Fatalf("binaryArraySearch() = %v; want %v", got, tc.result)
			}
		})
	}
}

func Benchmark_binaryArraySearch10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearch(slice, 344)
	}
}
func Benchmark_binaryArraySearch100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearch(slice, 344)
	}
}
func Benchmark_binaryArraySearch1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearch(slice, 344)
	}
}
func Benchmark_binaryArraySearch10K(t *testing.B) {
	slice := ints[0:10_000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearch(slice, 344)
	}
}
func Benchmark_binaryArraySearch100K(t *testing.B) {
	slice := ints[0:10_0000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearch(slice, 344)
	}
}
func Benchmark_binaryArraySearch1M(t *testing.B) {
	slice := ints[0:1_000_000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearch(slice, 344)
	}
}

// =====================================================================================================================
// binaryArraySearchIndex
// =====================================================================================================================

func Test_binaryArraySearchIndex(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		result int
	}{
		{[]int{2, 14, 15, 19, 26, 53, 58}, 53, 5},
		{[]int{2, 14, 15, 19, 26, 53, 58}, 17, -3},
		{[]int{-99, -60, -4, -1}, -60, 1},
		{[]int{-99, -60, -4, -1}, 60, -4},
		{[]int{-99, -60, -4, -1, 0, 30, 40, 99, 4999}, 30, 5},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got := binaryArraySearchIndex(tc.input, tc.target)
			if got != tc.result {
				t.Fatalf("binaryArraySearchIndex() = %v; want %v", got, tc.result)
			}
		})
	}
}

func Benchmark_binaryArraySearchIndex10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndex(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndex100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndex(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndex1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndex(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndex10K(t *testing.B) {
	slice := ints[0:10_000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndex(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndex100K(t *testing.B) {
	slice := ints[0:10_0000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndex(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndex1M(t *testing.B) {
	slice := ints[0:1_000_000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndex(slice, 344)
	}
}

// =====================================================================================================================
// binaryArraySearchIndexOptimized
// =====================================================================================================================

func Test_binaryArraySearchIndexOptimized(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		result int
	}{
		{[]int{2, 14, 15, 19, 26, 53, 58}, 53, 5},
		{[]int{2, 14, 15, 19, 26, 53, 58}, 17, -3},
		{[]int{-99, -60, -4, -1}, -60, 1},
		{[]int{-99, -60, -4, -1}, 60, -4},
		{[]int{-99, -60, -4, -1, 0, 30, 40, 99, 4999}, 30, 5},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got := binaryArraySearchIndexOptimized(tc.input, tc.target)
			if got != tc.result {
				t.Fatalf("binaryArraySearchIndexOptimized() = %v; want %v", got, tc.result)
			}
		})
	}
}

func Benchmark_binaryArraySearchIndexOptimized10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndexOptimized(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndexOptimized100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndexOptimized(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndexOptimized1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndexOptimized(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndexOptimized10K(t *testing.B) {
	slice := ints[0:10_000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndexOptimized(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndexOptimized100K(t *testing.B) {
	slice := ints[0:10_0000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndexOptimized(slice, 344)
	}
}
func Benchmark_binaryArraySearchIndexOptimized1M(t *testing.B) {
	slice := ints[0:1_000_000]
	t.ResetTimer()
	for i := 0; i< t.N; i ++ {
		binaryArraySearchIndexOptimized(slice, 344)
	}
}