package sum

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		list []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{4, 3, 2, 1}, 10},
		{[]int{1, -2, 3, 4}, 6},
		{[]int{-4, 4, 4, -4}, 0},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc), func(t *testing.T) {
			got := Sum(tc.list)
			if got != tc.want {
				t.Fatalf("Sum: got %v want %v", got, tc.want)
			}
		})
	}
}

func TestSumRecursively(t *testing.T) {
	tests := []struct {
		list []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{4, 3, 2, 1}, 10},
		{[]int{1, -2, 3, 4}, 6},
		{[]int{-4, 4, 4, -4}, 0},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc), func(t *testing.T) {
			got := SumRecursively(tc.list)
			if got != tc.want {
				t.Fatalf("Sum: got %v want %v", got, tc.want)
			}
		})
	}
}
