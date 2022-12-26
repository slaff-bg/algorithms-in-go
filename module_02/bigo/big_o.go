package bigo

import (
	"fmt"
	"strings"
)

// All of the code in this file is used when explaining Big O and viewing
// benchmarks to help illustrate a few points.

// O(1)
// Add is O(1) - constant time.
func Add(a, b int) int {
	for i := 0; i < 1000; i++ {
		// do something here ...
	}
	return a + b
}

// O(N)
// SumToMax is O(N) where N is the value of max (slow).
func SumToMax(max int) int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	return sum
}

// O(1)
// SumToMaxV2 illustrates that some functions can do the same thing
// but much faster. V2 is O(1) - constant time.
func SumToMaxV2(max int) int {
	return max * (max + 1) / 2
}

// O(N)
// SumVals is O(N) where N is the size of the vals slice.
// i.e. N == len(vals)
func SumVals(vals []int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

// O(N)
// Find has two inputs, but only one affects Big O.
// This is O(N) where N is the size of the list.
func Find(list []int, x int) int {
	for i, v := range list {
		if v == x {
			return i
		}
	}
	return -1
}

// O(XY)
// An example of a function that's Big O is determined by two inputs: x, y.
// Grid is O(XY) where X and Y are the values of x and y.
func Grid(x, y int) string {
	var sb strings.Builder
	chars := []byte{'x', 'o'}
	charIdx := 0

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			sb.WriteByte(chars[charIdx])
			charIdx = 1 - charIdx // alternates between 0 and 1
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

// O(N*M)
// PrintList has a big O of O(N*M) where N == value of n, and M == len(list)
func PrintList(list []int, n int) {
	for i := 0; i < n; i++ {
		for _, v := range list {
			fmt.Print(v)
		}
		fmt.Println()
	}
}

// O(N^3)
// Cube is O(N^3)
func Cube(n int) string {
	var sb strings.Builder
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			for z := 0; z < n; z++ {
				sb.WriteString(fmt.Sprintf("(%d, %d, %d)\n", x, y, z))
			}
		}
	}
	return sb.String()
}
