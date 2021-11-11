package main

import (
	"fmt"
	"testing"
)

// =====================================================================================================================
func Test_largestTwo(t *testing.T) {
	tests := []struct {
		input []int
		max   int
		second int
	}{
		{[]int{99, 29, 4, 28, -3, 23, 332}, 332, 99},
		{[]int{-3, -99, -494432, -78, -2222}, -3, -78},
		{[]int{3, 43, 74, 28, -24, 65, 2}, 74, 65},
		{[]int{959, 5432, 54, 628, -2613, 487, 395}, 5432, 959},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got, got2 := largestTwo(tc.input...)
			if got != tc.max {
				t.Fatalf("largestTwo() = %v; want max %v", got, tc.max)
			}
			if got2 != tc.second {
				t.Fatalf("largestTwo() = %v; want second %v", got2, tc.second)
			}
		})
	}
}
func Benchmark_largestTwo_10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largestTwo(slice...)
	}
}
func Benchmark_largestTwo_100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largestTwo(slice...)
	}
}
func Benchmark_largestTwo_1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largestTwo(slice...)
	}
}
func Benchmark_largestTwo_10K(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largestTwo(slice...)
	}
}
func Benchmark_largestTwo_100K(t *testing.B) {
	slice := ints[0:100000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largestTwo(slice...)
	}
}
func Benchmark_largestTwo_1000K(t *testing.B) {
	slice := ints[0:1000000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largestTwo(slice...)
	}
}

// =====================================================================================================================
func Test_sortingTwo(t *testing.T) {
	tests := []struct {
		input []int
		max   int
		second int
	}{
		{[]int{99, 29, 4, 28, -3, 23, 332}, 332, 99},
		{[]int{-3, -99, -494432, -78, -2222}, -3, -78},
		{[]int{3, 43, 74, 28, -24, 65, 2}, 74, 65},
		{[]int{959, 5432, 54, 628, -2613, 487, 395}, 5432, 959},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got, got2 := sortingTwo(tc.input...)
			if got != tc.max {
				t.Fatalf("sortingTwo() = %v; want max %v", got, tc.max)
			}
			if got2 != tc.second {
				t.Fatalf("sortingTwo() = %v; want second %v", got2, tc.second)
			}
		})
	}
}
func Benchmark_sortingTwo_10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		sortingTwo(slice...)
	}
}
func Benchmark_sortingTwo_100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		sortingTwo(slice...)
	}
}
func Benchmark_sortingTwo_1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		sortingTwo(slice...)
	}
}
func Benchmark_sortingTwo_10K(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		sortingTwo(slice...)
	}
}
func Benchmark_sortingTwo_100K(t *testing.B) {
	slice := ints[0:100000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		sortingTwo(slice...)
	}
}
func Benchmark_sortingTwo_1000K(t *testing.B) {
	slice := ints[0:1000000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		sortingTwo(slice...)
	}
}

// =====================================================================================================================
func Test_doubleTwo(t *testing.T) {
	tests := []struct {
		input []int
		max   int
		second int
	}{
		{[]int{99, 29, 4, 28, -3, 23, 332}, 332, 99},
		{[]int{-3, -99, -494432, -78, -2222}, -3, -78},
		{[]int{3, 43, 74, 28, -24, 65, 2}, 74, 65},
		{[]int{959, 5432, 54, 628, -2613, 487, 395}, 5432, 959},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got, got2 := doubleTwo(tc.input...)
			if got != tc.max {
				t.Fatalf("doubleTwo() = %v; want max %v", got, tc.max)
			}
			if got2 != tc.second {
				t.Fatalf("doubleTwo() = %v; want second %v", got2, tc.second)
			}
		})
	}
}
func Benchmark_doubleTwo_10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		doubleTwo(slice...)
	}
}
func Benchmark_doubleTwo_100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		doubleTwo(slice...)
	}
}
func Benchmark_doubleTwo_1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		doubleTwo(slice...)
	}
}
func Benchmark_doubleTwo_10K(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		doubleTwo(slice...)
	}
}
func Benchmark_doubleTwo_1000K(t *testing.B) {
	slice := ints[0:1000000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		doubleTwo(slice...)
	}
}

// =====================================================================================================================
func Test_mutableTwo(t *testing.T) {
	tests := []struct {
		input []int
		max   int
		second int
	}{
		{[]int{99, 29, 4, 28, -3, 23, 332}, 332, 99},
		{[]int{-3, -99, -494432, -78, -2222}, -3, -78},
		{[]int{3, 43, 74, 28, -24, 65, 2}, 74, 65},
		{[]int{959, 5432, 54, 628, 5, -2613, 487, 395}, 5432, 959},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got, got2 := mutableTwo(tc.input...)
			if got != tc.max {
				t.Fatalf("mutableTwo() = %v; want max %v", got, tc.max)
			}
			if got2 != tc.second {
				t.Fatalf("mutableTwo() = %v; want second %v", got2, tc.second)
			}
		})
	}
}
func Benchmark_mutableTwo_10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mutableTwo(slice...)
	}
}
func Benchmark_mutableTwo_100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mutableTwo(slice...)
	}
}
func Benchmark_mutableTwo_1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mutableTwo(slice...)
	}
}
func Benchmark_mutableTwo_10K(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mutableTwo(slice...)
	}
}
func Benchmark_mutableTwo_100K(t *testing.B) {
	slice := ints[0:100000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mutableTwo(slice...)
	}
}
func Benchmark_mutableTwo_1000K(t *testing.B) {
	slice := ints[0:1000000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mutableTwo(slice...)
	}
}

// =====================================================================================================================
func Test_tournamentTwo(t *testing.T) {
	tests := []struct {
		input []int
		max   int
		second int
	}{
		{[]int{99, 29, 4, 28, -3, 23, 5, 332}, 332, 99},
		{[]int{-3, -99, -494432, -78, -555, -2222}, -3, -78},
		{[]int{3, 43, 74, 28, -24, 65, 5, 2}, 74, 65},
		{[]int{959, 5432, 54, 628, 5, -2613, 487, 395}, 5432, 959},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got, got2 := tournamentTwo(tc.input...)
			if got != tc.max {
				t.Fatalf("tournamentTwo() = %v; want max %v", got, tc.max)
			}
			if got2 != tc.second {
				t.Fatalf("tournamentTwo() = %v; want second %v", got2, tc.second)
			}
		})
	}
}
func Benchmark_tournamentTwo_10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		tournamentTwo(slice...)
	}
}
func Benchmark_tournamentTwo_100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		tournamentTwo(slice...)
	}
}
func Benchmark_tournamentTwo_1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		tournamentTwo(slice...)
	}
}
func Benchmark_tournamentTwo_10K(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		tournamentTwo(slice...)
	}
}
func Benchmark_tournamentTwo_100K(t *testing.B) {
	slice := ints[0:100000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		tournamentTwo(slice...)
	}
}
func Benchmark_tournamentTwo_1000K(t *testing.B) {
	slice := ints[0:1000000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		tournamentTwo(slice...)
	}
}