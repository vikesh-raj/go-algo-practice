package dynamicprogramming

// https://en.wikipedia.org/wiki/Composition_(combinatorics)
//
// There are n steps in a staircase
// You can climb upto k steps at a time.
// Find different combinations for climbing the staircase

// Eg: n = 5 , k = 2
// 1+1+1+1+1, 1+1+1+2, 1+1+2+1, 1+2+1+1, 2+1+1+1, 2+2+1, 2+1+2, 1+2+2
// Total of 8 ways

// If there are no restrictions of number of steps,
// the total combinations are 2^(n-1)

// If k == 2, It follows fibonacci series with 1, 2 as starting numbers
// F(n) = F(n-1) + F(n-2)
// F(1) = 1
// F(2) = 2

// F(3) = 1 + 2 = 3
// F(4) = 3 + 2 = 5
// F(5) = 5 + 3 = 8
// F(6) = 8 + 5 = 13
// F(7) = 13 + 8 = 21
// F(8) = 21 + 13 = 34

// We can now see that for k == 3, it follows tribonacci series
// T(n) = T(n-1) + T(n-2) + T(n-3)
// T(1) = 1
// T(2) = 2
// T(3) = 4
// T(4) = 7
// T(5) = 13
// T(6) = 24

// For general case,

// S(n) =  sum(1 to k) S(n - k) if n > k
// else S(n) = 2^(n-1) if n <= k

func StairCaseCombinations(totalSteps, maxSteps int) int {
	if totalSteps == 0 {
		return 0
	}

	sums := make([]int, totalSteps+1)
	sums[0] = 0
	sums[1] = 1

	for i := 2; i <= maxSteps && i <= totalSteps; i++ {
		sums[i] = 2 * sums[i-1]
	}

	for i := maxSteps + 1; i <= totalSteps; i++ {
		for j := 1; j <= maxSteps; j++ {
			sums[i] += sums[i-j]
		}
	}

	return sums[totalSteps]
}
