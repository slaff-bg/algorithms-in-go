package sum

import "fmt"

// Sum all the numbers in the list.
// = = = = = = = = = = = = = = = = = = = =

func Sum(list []int) int {
	total := 0
	for _, v := range list {
		total += v
	}
	return total
}

// Bad practice. Just for exercise.
func SumRecursively(list []int) int {
	fmt.Println(list)
	if len(list) == 0 {
		return 0
	}
	return list[0] + SumRecursively(list[1:])
}
