package stringreversing

import "strings"

// String reversing.
// = = = = = = = = = = = = = = = = = = = =

func Reverse(word string) string {
	var res string

	for i := len(word) - 1; i >= 0; i-- {
		res += string(word[i])
	}

	return res
}

// Best approach!
func ReverseUsingRunes(word string) string {
	var res string

	for _, r := range word {
		res = string(r) + res
	}

	return res
}

func ReverseAlternative(word string) string {
	var res string

	for i := 0; i < len(word); i++ {
		res = string(word[i]) + res
	}

	return res
}

func ReverseUsingStringsBilder(word string) string {
	var res strings.Builder

	for i := len(word) - 1; i >= 0; i-- {
		res.WriteByte(word[i])
	}

	return res.String()
}
