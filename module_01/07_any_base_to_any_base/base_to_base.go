package anybasetoanybase

const charset = "0123456789ABCDEFGHIJ"

func BaseToBase(value string, base, newBase int) string {
	return decimalToBase(baseToDecimal(value, base), newBase)
}

func baseToDecimal(value string, base int) int {
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

		res += index * multiplier
		multiplier *= base
	}

	return res
}

func decimalToBase(dec, base int) string {
	res := ""

	for dec > 0 {
		rem := dec % base
		dec = dec / base
		res = string(charset[rem]) + res
	}

	return res
}
