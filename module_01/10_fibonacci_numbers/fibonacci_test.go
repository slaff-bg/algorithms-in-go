package fibonaccinumbers

import (
	"fmt"
	"testing"
)

type tC struct {
	n, want int
}

var tests = []tC{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{14, 377},
	{22, 17711},
	{25, 75025},
	// These test cases may be much slower depending on the solution.
	// Feel free to comment it out if it is too slow.
	// {34, 5702887},
	// {38, 39088169},
	// {40, 102334155},
	// {41, 165580141},
	// ... or just comment out the test and benchmark functions
	// of the first (slowest) approach:
	// TestFibonacci() and BenchmarkFibonacci().
}

func TestFibonacci(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("n=%v", tc.n), func(t *testing.T) {
			got := Fibonacci(tc.n)
			if got != tc.want {
				t.Fatalf("Fibonacci: %v; want %v", got, tc.want)
			}
		})
	}
}

func TestAltFibonacci(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("n=%v", tc.n), func(t *testing.T) {
			got := AltFibonacci(tc.n)
			if got != tc.want {
				t.Fatalf("AltFibonacci: %v; want %v", got, tc.want)
			}
		})
	}
}

func TestSecondAltFibonacci(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("n=%v", tc.n), func(t *testing.T) {
			got := SecondAltFibonacci(tc.n)
			if got != tc.want {
				t.Fatalf("SecondAltFibonacci: %v; want %v", got, tc.want)
			}
		})
	}
}

// - - - - - - - - - - - - - - - - -
// BENCHMARKS
// - - - - - - - - - - - - - - - - -

var bmResF int

func BenchmarkFibonacci(b *testing.B) {
	var f int
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			b.Run("n=%v", func(b *testing.B) {
				// always record the result of Fib to prevent
				// the compiler eliminating the function call.
				f = Fibonacci(tc.n)
			})
		}
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	bmResF = f
}

var bmResAF int

func BenchmarkAltFibonacci(b *testing.B) {
	var f int
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			b.Run("n=%v", func(b *testing.B) {
				f = AltFibonacci(tc.n)
			})
		}
	}

	bmResAF = f
}

var bmResSAF int

func BenchmarkSecondAltFibonacci(b *testing.B) {
	var f int
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			b.Run("n=%v", func(b *testing.B) {
				f = SecondAltFibonacci(tc.n)
			})
		}
	}

	bmResSAF = f
}
