package dynamicprogramming

//
// https://en.wikipedia.org/wiki/Fibonacci_number
//
// F(n) = F(n - 1) + F(n - 2)
// F(0) = 0
// F(1) = 1

func Fibonacci(n int) int {
	sequence := make([]int, n+1)
	sequence[0] = 0
	sequence[1] = 1

	for i := 2; i <= n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	return sequence[n]
}
