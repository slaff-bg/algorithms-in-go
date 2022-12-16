package findtwothatsum

// Find two numbers in a list that sum to a given amount
//
// FindTwoThatSum will look for two numbers in the provided
// numbers slice that sum up to the sum argument. It will then
// return the indices of each of these numbers.
//
// If there are multiple correct answers, any of them may be
// used. If there are no correct answers, (-1, -1) is returned.
//
// Eg:
//
//	FindTwoThatSum([1,2,3,4], 7) => (2, 3)
//	FindTwoThatSum([0, -1, 1], 0) => (1, 2)
//	FindTwoThatSum([0, 1, 1], 0) => (-1, -1)
//
// Or if we have duplicate answers any of them are okay. All
// of the following are correct.
//
//	FindTwoThatSum([10, 1, 12, 3, 7, 2, 2, 1], 4) => (5, 6)
//	FindTwoThatSum([10, 1, 12, 3, 7, 2, 2, 1], 4) => (1, 3)
//	FindTwoThatSum([10, 1, 12, 3, 7, 2, 2, 1], 4) => (1, 7)
//
// Each number will only be used once, and the original numbers
// list should not be rearranged or altered in any way.
//
// If you want to challenge yourself further, try adding different
// criteria for which numbers will be used when there are multiple
// correct answers. Take a look at the *_ext.go file
func FindTwoThatSum(numbers []int, sum int) (int, int) {
	for i, v := range numbers {
		for j, v2 := range numbers {
			if i == j {
				continue
			}

			if v+v2 == sum {
				return i, j
			}
		}
	}

	return -1, -1
}

type resIndexes struct {
	res1, res2 int
}

var possibleRes = []resIndexes{}

func FindTwoThatSumAlt(numbers []int, sum int) (int, int) {
	for i, vo := range numbers {
		for j := len(numbers); j > i; j-- {
			if vo+numbers[j] == sum {
				possibleRes = append(possibleRes, resIndexes{i, j})
			}
		}
	}

	if len(possibleRes) > 0 {
		defer func() {
			possibleRes = nil
		}()
		// @todo: implement conditions here ...

		return possibleRes[0].res1, possibleRes[0].res2
	}

	return -1, -1
}

func FindTwoThatSumByCriteria(numbers []int, sum int, criteria string) (int, int) {
	for i, vo := range numbers {
		for j := len(numbers); j > i; j-- {
			if vo+numbers[j] == sum {
				possibleRes = append(possibleRes, resIndexes{i, j})
			}
		}
	}

	if len(possibleRes) > 0 {
		defer func() {
			possibleRes = nil
		}()
		// @todo: implement conditions here ...

		return possibleRes[0].res1, possibleRes[0].res2
	}

	return -1, -1
}
