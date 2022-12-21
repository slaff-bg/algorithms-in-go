package greatestcommondivisior

// Using recursion ...
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Using recursion ...
func GCDRecursionAlt(a, b int) int {
	if b == 0 {
		return a
	}

	return GCDRecursionAlt(b, a%b)
}

// Without using division ...
func GCDNoDivisionAlt(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}

	return a
}

// Remainder division (the slowest approach) ...
func GCDRemainderDivAlt(a, b int) int {
	// Step 1. to be sure that the first variable contains a lower value
	if a >= b {
		a, b = b, a
	}

	// Step 2. no additional operations are required if the higher value
	// can be divided by the smaller value
	if b%a == 0 {
		return a
	}

	// Step 3. rotates only the values that are below half of the smaller number
	for i := int(a / 2); i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}

	return 1
}
