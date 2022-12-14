package findtwothatsum

// The task requirement from the find_two_that_sum.go file,
// plus a challenge ...
//
// If you want to challenge yourself further, try adding
// different criteria for which numbers will be used when
// there are multiple correct answers.
// Eg:
//  1. Always return the lowest index possible for the first value.
//  2. Always return the index of lowest value possible for the
//     first return value.
//  3. Always return the index of the two values who have a minimal
//     difference. Eg prefer the values 2, 2 over 1, 3 over 0, 4
//     for the sum of 4.

type resIndVal struct {
	i1, i2, v1, v2 int
}

var possibleResXt = []resIndVal{}

// criteria:
//   - LI (Always return the lowest index possible for the first value.)
//   - LV (Always return the index of lowest value possible for the
//     first return value.)
//   - MinD (Always return the index of the two values who have a minimal
//     difference.
//
// E.g. (MinD)
// prefer the values 2, 2 over 1, 3 over 0, 4 for the sum of 4.
func FindTwoThatSumByCriteriaXt(numbers []int, sum int, criteria string) (int, int) {
	clnPossibleResXt := possibleResXt
	for i, vo := range numbers {
		for j := len(numbers) - 1; j > i; j-- {
			if vo+numbers[j] == sum {
				clnPossibleResXt = append(clnPossibleResXt, resIndVal{i, j, vo, numbers[j]})
			}
		}
	}

	if len(clnPossibleResXt) > 0 {
		defer func() {
			clnPossibleResXt = nil
		}()

		switch criteria {
		case "LI":
			return clnPossibleResXt[0].i1, clnPossibleResXt[0].i2
		case "LV":
			return getByLV(clnPossibleResXt)
		case "MinD":
			return getByMinD(clnPossibleResXt)
		}
	}

	return -1, -1
}

// Returns the index of lowest value possible for the first return value.
func getByLV(clnPossibleResXt []resIndVal) (int, int) {
	if len(clnPossibleResXt) == 1 {
		return clnPossibleResXt[0].i1, clnPossibleResXt[0].i2
	}

	// fI => keeps the loop Index
	fI, lowest := 0, clnPossibleResXt[0].v1
	for i := 1; i < len(clnPossibleResXt); i++ {
		if lowest > clnPossibleResXt[i].v1 {
			fI, lowest = i, clnPossibleResXt[i].v1
		}
	}

	return clnPossibleResXt[fI].i1, clnPossibleResXt[fI].i2
}

// Returns the index of the two values who have a minimal difference.
// e.g.
// prefer the values 2, 2 over 1, 3 over 0, 4 for the sum of 4.
func getByMinD(clnPossibleResXt []resIndVal) (int, int) {
	if len(clnPossibleResXt) == 1 {
		return clnPossibleResXt[0].i1, clnPossibleResXt[0].i2
	}

	// fI => keeps the loop Index
	fI, diff := 0, checkDiffLength(clnPossibleResXt[0].v1, clnPossibleResXt[0].v2)
	for i := 1; i < len(clnPossibleResXt); i++ {
		tmpDiff := checkDiffLength(clnPossibleResXt[i].v1, clnPossibleResXt[i].v2)
		if diff > tmpDiff {
			fI, diff = i, tmpDiff
		}
	}

	return clnPossibleResXt[fI].i1, clnPossibleResXt[fI].i2
}

func checkDiffLength(v1, v2 int) int {
	if v1 < v2 {
		return v2 - v1
	}

	return v1 - v2
}
