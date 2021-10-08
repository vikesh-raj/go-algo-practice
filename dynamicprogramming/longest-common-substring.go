package dynamicprogramming

// Longest common substring

// Given two strings s and t, find the longest substring (contiguous characters)
// that is both in s and t.

// Example: s = "tofoodie", t = "toody", longest common substring = "ood"

// Formula
// D[i,j] = 0  if s[i] != t[j]
//        = 1 + D[i-1][j-1] if  s[i] == t[j]

func LongestCommonSubstring(s, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	srunes := []rune(s)
	trunes := []rune(t)

	d := make([][]int, len(srunes)+1)
	for i := 0; i <= len(srunes); i++ {
		d[i] = make([]int, len(trunes)+1)
	}

	max := 0
	rmaxIndex := 0

	for i := 1; i <= len(srunes); i++ {
		for j := 1; j <= len(trunes); j++ {
			if srunes[i-1] == trunes[j-1] {
				d[i][j] = 1 + d[i-1][j-1]
				if d[i][j] > max {
					max = d[i][j]
					rmaxIndex = i
				}
			}
		}
	}

	output := make([]rune, max)
	for i := len(output) - 1; i >= 0; i-- {
		output[i] = srunes[rmaxIndex-1]
		rmaxIndex--
	}
	return string(output)
}
