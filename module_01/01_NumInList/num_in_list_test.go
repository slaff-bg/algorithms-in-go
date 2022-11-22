package numinlist

import (
	"fmt"
	"testing"
)

func TestNumInList(t *testing.T) {
	tests := []struct {
		list []int
		num  int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, 4, true},
		{[]int{1, 2, 3, 4, 5}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 0, false},
		{[]int{1, 2, 3, 4, 5}, 6, false},
		{[]int{1, 2, 3, 4, 5}, -4, false},
		{[]int{1, 2, 3, -4, 5}, -4, true},
		{[]int{4, 4, 4, 4, 4}, 4, true},
		{[]int{4, 4, 4, 4, 4}, -4, false},
		{[]int{-4, -4, -4, -4, -4}, 4, false},
		{[]int{1, 1, 1, 1, 1}, 4, false},
		{[]int{8, 14, 9, 21, 7}, 7, true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v,%v]", tc.list, tc.num), func(t *testing.T) {
			got := NumInList(tc.list, tc.num)
			if got != tc.want {
				t.Fatalf("NumInList: got %v want %v", got, tc.want)
			}
		})
	}
}
