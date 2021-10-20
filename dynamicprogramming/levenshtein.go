package dynamicprogramming

import (
	"github.com/vikesh-raj/go-algo-practice/helpers"
)

// Levenshtein distance

// Find the Levenshtein distance between two strings.
// Example: s = "kitten", t = "sitting", levenshtein distance = 3
// kitten -> sitten  (substitute k with s)
// sitten -> sittin  (substitute e with i)
// sittin -> sitting (add g in the end)

// Formula
// D[i,j] = 1 + min(D[i-1,j], D[i,j-1], D[i-1,j-1])  if s[i] != t[j]
//        = D[i-1][j-1] if  s[i] == t[j]

func LevenshteinDistance(s, t string) int {
	srunes := []rune(s)
	trunes := []rune(t)

	if len(srunes) == 0 || len(trunes) == 0 {
		return helpers.Max(len(srunes), len(trunes))
	}

	d := make([][]int, len(srunes)+1)
	for i := 0; i <= len(srunes); i++ {
		d[i] = make([]int, len(trunes)+1)
	}

	for i := 1; i <= len(srunes); i++ {
		for j := 1; j <= len(trunes); j++ {
			if srunes[i-1] == trunes[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = 1 + helpers.MinMulti(d[i-1][j], d[i][j-1], d[i-1][j-1])
			}
		}
	}

	return d[len(srunes)][len(trunes)]
}
