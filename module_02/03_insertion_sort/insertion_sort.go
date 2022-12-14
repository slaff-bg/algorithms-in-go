package insertionsort

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	// sort.Ints(list)
	// sort.IntSlice(list).Sort()
	var sorted []int
	for _, v := range list {
		sorted = insert(sorted, v)
	}
	copy(list, sorted)
}

func insert(dest []int, val int) []int {
	for i, v := range dest {
		if val < v {
			return append(dest[:i], append([]int{val}, dest[i:]...)...)
		}
	}
	return append(dest, val)
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
// func InsertionSortString(list []string) {
// }

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
// func InsertionSortPerson(people []Person) {
// }

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
// func InsertionSort(list sort.Interface) {
// }
