package fibonaccinumbers

/*
4
-------------
3+2 		= 3
	(3)2+1 		= 2
		(2)1+0 		= 1
		(1)1 		= 1
	(2)1+0 		= 1
		(1)1 		= 1
		(0)0 		= 0
*/

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func AltFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	nMinus2 := 0
	nMinus1 := 1
	var cur int
	for i := 2; i <= n; i++ {
		cur = nMinus2 + nMinus1
		nMinus2 = nMinus1
		nMinus1 = cur
	}

	return cur
}

// n = n-1 + n-2
func SecondAltFibonacci(times int) int {
	if times < 2 {
		return times
	}

	var res int
	cur, prev := 1, 0

	for i := 0; i < times-1; i++ {
		res = cur + prev
		prev = cur
		cur = res
	}

	return res
}
