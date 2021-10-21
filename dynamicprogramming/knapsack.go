package dynamicprogramming

import (
	"fmt"

	"github.com/vikesh-raj/go-algo-practice/helpers"
)

// Knapsack problem

// Given n items, each with a value and weight and knapsack capable of
// supporting weight W, find the largest-valued subset possible to carry in the knapsack.

// Formula
// D[i,j] = D[i-1, j] if j < wi
//        = max(D[i-1, j], D[i-1,j-wi] + vi)

type Item struct {
	Value  int
	Weight int
}

func Knapsack(items []Item, capacity int) (int, []int) {

	d := make([][]int, len(items)+1)
	for i := 0; i <= len(items); i++ {
		d[i] = make([]int, capacity+1)
	}

	for i := 1; i <= len(items); i++ {
		for j := 1; j <= capacity; j++ {
			if j < items[i-1].Weight {
				d[i][j] = d[i-1][j]
			} else {
				d[i][j] = helpers.Max(d[i-1][j], d[i-1][j-items[i-1].Weight]+items[i-1].Value)
			}
		}
	}

	fmt.Println(d)
	maxValue := d[len(items)][capacity]
	output := make([]int, 0)

	i := len(items)
	c := capacity
	for i > 0 && c > 0 {
		if d[i][c] == d[i-1][c] {
			i--
		} else {
			output = append(output, i-1)
			c -= items[i-1].Weight
			i--
		}
	}

	return maxValue, reverse(output)
}

func reverse(a []int) []int {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
	return a
}
