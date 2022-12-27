package bubblesort

import (
	"sort"
)

// O(N^2)
// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {

	// Basic approach.
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

	/*
		// Optimized approach.
		for i := 0; i < len(list); i++ {
			check := true
			for j := 0; j < len(list)-1-i; j++ {
				if list[j] > list[j+1] {
					list[j], list[j+1] = list[j+1], list[j]
					check = false
				}
			}
			if check {
				break
			}
		}
	*/

	/*
		// A better alternative approach even without break directive (fewer permutations; no second check).
		for i := len(list) - 1; i >= 0; i-- {
			for j := i; j < len(list)-1; j++ {
				if list[j] > list[j+1] {
					list[j], list[j+1] = list[j+1], list[j]
					continue
				}

				break
			}
		}
	*/
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
// func BubbleSortPerson(people []Person) {
// }

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
}
