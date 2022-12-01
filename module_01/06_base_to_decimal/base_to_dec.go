package basetodecimal

import "fmt"

func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEFGHIJ"
	multiplier := 1
	res := 0

	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for j, char := range charset {
			if rune(value[i]) == char { // value[i] == byte(char)
				index = j
				break
			}
		}

		// Shouldn't ever happen ...
		// if index < 0 {
		// 	panic("Something went wrong!")
		// }

		res += index * multiplier
		multiplier *= base
	}

	return res
}

func BaseToDecAlt(value string, base int) int {
	res := 0
	multiplier := 1

	for i := len(value) - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(value[i]), "%X", &val)
		res += val * multiplier
		multiplier *= base
	}

	return res
}
