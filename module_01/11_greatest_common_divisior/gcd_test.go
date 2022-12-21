package greatestcommondivisior

import (
	"fmt"
	"testing"
)

type tC struct {
	a, b, want int
}

var tests = []tC{
	{10, 5, 5},
	{23, 5, 1},
	{25, 5, 5},
	{30, 15, 15},
	{30, 9, 3},
	{100, 9, 1},
	// Sometimes it is easier to produce test cases by
	// writing out how each value is derived. In this case
	// numbers are composed of prime factors and the
	// common prime factors should be present in the answer
	// so writing test cases this way makes them easier to
	// verify visually.
	{
		2 * 2 * 3 * 3 * 5,
		2 * 3 * 5 * 7 * 13,
		2 * 3 * 5,
	}, {
		2 * 2 * 3 * 3 * 13,
		2 * 3 * 5 * 7 * 13,
		2 * 3 * 13,
	}, {
		2 * 3 * 5 * 7 * 11 * 13 * 17 * 19,
		3 * 3 * 7 * 7 * 11 * 11 * 17 * 17,
		3 * 7 * 11 * 17,
	},
}

func TestGCD(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v,%v)", tc.a, tc.b), func(t *testing.T) {
			got := GCD(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("GCD: %v; want %v", got, tc.want)
			}
		})
	}
}

func TestGCDRecursionAlt(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v,%v)", tc.a, tc.b), func(t *testing.T) {
			got := GCDRecursionAlt(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("GCDRecursionAlt: %v; want %v", got, tc.want)
			}
		})
	}
}

func TestGCDNoDivisionAlt(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v,%v)", tc.a, tc.b), func(t *testing.T) {
			got := GCDNoDivisionAlt(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("GCDNoDivisionAlt: %v; want %v", got, tc.want)
			}
		})
	}
}

func TestGCDRemainderDivAlt(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v,%v)", tc.a, tc.b), func(t *testing.T) {
			got := GCDRemainderDivAlt(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("GCDRemainderDivAlt: %v; want %v", got, tc.want)
			}
		})
	}
}

// go test -run=TestGCD
// go test -run=TestGCDRecursionAlt
// go test -run=TestGCDNoDivisionAlt
// go test -run=TestGCDRemainderDivAlt

// - - - - - - - - - - - - - - - - -
// BENCHMARKS
// - - - - - - - - - - - - - - - - -

func BenchmarkGCDRecursionAlt(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			b.Run("n=%v", func(b *testing.B) {
				GCDRecursionAlt(tc.a, tc.b)
			})
		}
	}
}

func BenchmarkGCDNoDivisionAlt(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			b.Run("n=%v", func(b *testing.B) {
				GCDNoDivisionAlt(tc.a, tc.b)
			})
		}
	}
}

// go test -bench=BenchmarkGCDRecursionAlt
// go test -bench=BenchmarkGCDNoDivisionAlt
