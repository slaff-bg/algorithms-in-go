package findtwothatsum

import (
	"fmt"
	"testing"
)

func TestFindTwoThatSumByCriteriaXt(t *testing.T) {
	type cRes struct {
		criteria   string
		res1, res2 int
	}

	tests := []struct {
		numbers []int
		sum     int
		cr      []cRes
	}{
		{[]int{10, 1, 12, 3, 7, 2, 2}, 4, []cRes{
			{"LI", 1, 3},
			{"LV", 1, 3},
			{"MinD", 5, 6},
		}},
		{[]int{10, 2, 2, 1, 12, 3, 7}, 4, []cRes{
			{"LI", 1, 2},
			{"LV", 3, 5},
			{"MinD", 1, 2},
		}},
		{[]int{2, -2, 2, -6, 10, 6, -1, 5}, 4, []cRes{
			{"LI", 0, 2},
			{"LV", 3, 4},
			{"MinD", 0, 2},
		}},
		{[]int{-2, 2, -6, 10, 6, -1, 5}, 4, []cRes{
			{"LI", 0, 4},
			{"LV", 2, 3},
			{"MinD", 5, 6},
		}},
		{[]int{10, 1, 12, 3, 7, 2, 2, -1, 5}, 4, []cRes{
			{"LI", 1, 3},
			{"LV", 7, 8},
			{"MinD", 5, 6},
		}},
		{[]int{-1, 5, 10, 1, 12, 3, 7, 2}, 4, []cRes{
			{"LI", 0, 1},
			{"LV", 0, 1},
			{"MinD", 3, 5},
		}},
	}

	for _, tc := range tests {
		for _, supTc := range tc.cr {
			t.Run(fmt.Sprintf("numbers: %v, sum: %v, criteria: %v", tc.numbers, tc.sum, supTc.criteria), func(t *testing.T) {
				i, j := FindTwoThatSumByCriteriaXt(tc.numbers, tc.sum, supTc.criteria)
				got := tc.numbers[i] + tc.numbers[j]
				if got != tc.sum {
					t.Fatalf("FindTwoThatSumByCriteriaXt: (%v, %v); sum = %v; want sum %v", i, j, got, tc.sum)
				}

				if i != supTc.res1 || j != supTc.res2 {
					t.Fatalf("FindTwoThatSumByCriteriaXt: criteria(%v); got indices (%v, %v); want indices (%v, %v)", supTc.criteria, i, j, supTc.res1, supTc.res2)
				}
			})
		}
	}
}
