package helpers

import (
	"math"
)

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxArray(array []int) int {
	max := math.MinInt32
	for _, value := range array {
		if value > max {
			max = value
		}
	}
	return max
}

func MinArray(array []int) int {
	min := math.MaxInt32
	for _, value := range array {
		if value < min {
			min = value
		}
	}
	return min
}

func MaxMulti(values ...int) int {
	return MaxArray(values)
}

func MinMulti(values ...int) int {
	return MinArray(values)
}

func MaxArrayWithIndex(array []int) (int, int) {
	max := math.MinInt32
	maxIndex := -1
	for i, value := range array {
		if value > max {
			max = value
			maxIndex = i
		}
	}
	return max, maxIndex
}

func MinArrayWithIndex(array []int) (int, int) {
	min := math.MaxInt32
	minIndex := -1
	for i, value := range array {
		if value < min {
			min = value
			minIndex = i
		}
	}
	return min, minIndex
}
