package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var ints []int

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	ints = make([]int, 1000000)

	for i := 0; i < 1000000; i++ {
		ints[i] = rand.Int()
	}

	m.Run()
}

func Test_largest(t *testing.T) {
	tests := []struct {
		input []int
		Max   int
		Index int
	}{
		{[]int{99, 29, 4, 28, -3, 23, 332}, 332, 6},
		{[]int{-3, -99, -494432, -99, -2222}, -3, 0},
		{[]int{3, 43, 74, 28 - 24, 65, 2}, 74, 2},
		{[]int{959, 5432, 54, 628, -2613, 487, 395}, 5432, 1},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got, got1, err := largest(tc.input...)
			if got != tc.Max {
				t.Fatalf("largest() = %v; want max %v", got, tc.Max)
			}
			if got1 != tc.Index {
				t.Fatalf("largest() = %v; want index %v", got, tc.Index)
			}
			require.NoError(t, err)
		})
	}
	t.Run(fmt.Sprintf("empty slice: %v", []int{}), func(*testing.T) {
		got, got1, err := largest([]int{}...)
		if got != 0 {
			t.Fatalf("largest() = %v; want max %v", got, 0)
		}
		if got1 != 0 {
			t.Fatalf("largest() = %v; want index %v", got, 0)
		}
		require.Error(t, err)
	})

}

func Benchmark_largest10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largest(slice...)
	}
}
func Benchmark_largest100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largest(slice...)
	}
}
func Benchmark_largest1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largest(slice...)
	}
}
func Benchmark_largest10K(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largest(slice...)
	}
}
func Benchmark_largest100K(t *testing.B) {
	slice := ints[0:100000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largest(slice...)
	}
}
func Benchmark_largest1M(t *testing.B) {
	slice := ints[0:1000000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		largest(slice...)
	}
}

// =====================================================================================================================
func Test_alternate(t *testing.T) {
	tests := []struct {
		input []int
		Max   int
	}{
		{[]int{99, 29, 4, 28, -3, 23, 332}, 332},
		{[]int{-3, -99, -494432, -99, -2222}, -3},
		{[]int{3, 43, 74, 28 - 24, 65, 2}, 74},
		{[]int{959, 5432, 54, 628, -2613, 487, 395}, 5432},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("slice: %v", tc.input), func(*testing.T) {
			got := alternate(tc.input...)
			if got != tc.Max {
				t.Fatalf("alternate() = %v; want max %v", got, tc.Max)
			}
		})
	}
}
func Benchmark_alternate10(t *testing.B) {
	slice := ints[0:10]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		alternate(slice...)
	}
}
func Benchmark_alternate100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		alternate(slice...)
	}
}
func Benchmark_alternate1K(t *testing.B) {
	slice := ints[0:1000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		alternate(slice...)
	}
}
func Benchmark_alternate10K(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		alternate(slice...)
	}
}
func Benchmark_alternate100K(t *testing.B) {
	slice := ints[0:100000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		alternate(slice...)
	}
}
func Benchmark_alternate1M(t *testing.B) {
	slice := ints[0:1000000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		alternate(slice...)
	}
}
