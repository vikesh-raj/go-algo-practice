package dynamicprogramming

import "github.com/vikesh-raj/go-algo-practice/helpers"

// Longest common subsequence

// Given two strings s and t, find the longest common subsequence

// Example: s = "xabayycczz", t = "pabqaqcqcr", longest common subsequence = "abacc"

// Formula
// D[i,j] = max(D[i-1,j], D[i,j-1])  if s[i] != t[j]
//        = 1 + D[i-1][j-1] if  s[i] == t[j]

func LongestCommonSubsequence(s, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	srunes := []rune(s)
	trunes := []rune(t)

	d := make([][]int, len(srunes)+1)
	for i := 0; i <= len(srunes); i++ {
		d[i] = make([]int, len(trunes)+1)
	}

	for i := 1; i <= len(srunes); i++ {
		for j := 1; j <= len(trunes); j++ {
			if srunes[i-1] == trunes[j-1] {
				d[i][j] = 1 + d[i-1][j-1]
			} else {
				d[i][j] = helpers.Max(d[i-1][j], d[i][j-1])
			}
		}
	}

	maxSeqLen := d[len(srunes)][len(trunes)]
	output := make([]rune, maxSeqLen)

	i := len(srunes)
	j := len(trunes)
	k := maxSeqLen - 1
	for i > 0 && j > 0 {
		if d[i][j] == d[i-1][j] {
			i--
		} else if d[i][j] == d[i][j-1] {
			j--
		} else {
			output[k] = trunes[j-1]
			k--
			i--
			j--
		}
	}
	return string(output)
}
