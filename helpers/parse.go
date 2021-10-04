package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseArrayOfInt(s string) ([]int, error) {
	words := strings.Fields(s)
	output := make([]int, 0, len(words))
	for _, word := range words {
		num, err := strconv.ParseInt(word, 10, 32)
		if err != nil {
			return nil, err
		}
		output = append(output, int(num))
	}
	return output, nil
}

func ParseArrayOfFloat(s string) ([]float64, error) {
	words := strings.Fields(s)
	output := make([]float64, 0, len(words))
	for _, word := range words {
		num, err := strconv.ParseFloat(word, 64)
		if err != nil {
			return nil, err
		}
		output = append(output, num)
	}
	return output, nil
}

func ParseMatrixInt(rows []string) ([][]int, error) {
	output := make([][]int, 0, len(rows))
	for i, row := range rows {
		numsArray, err := ParseArrayOfInt(row)
		if err != nil {
			return nil, fmt.Errorf("unable to parse row %d : (%s), error %v", i, row, err)
		}
		output = append(output, numsArray)
	}
	return output, nil
}
