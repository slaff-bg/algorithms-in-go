package decimaltoanotherbase

// Returns a string representing the provided decimal
// number in the provided base. This is limited to
// 2-16 for simplicity.
// e.g.
// DecToBase(14, 16) 	=> "E"
// DecToBase(14, 2) 	=> "1110"
func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEFGHIJ"
	res := ""

	for dec > 0 {
		rem := dec % base
		dec = dec / base
		res = string(charset[rem]) + res
		// res = fmt.Sprintf("%X%s", rem, res)
	}

	return res
}
