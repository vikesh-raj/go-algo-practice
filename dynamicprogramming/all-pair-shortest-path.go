package dynamicprogramming

import "math"

// Floyd Warshall Algorithm
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm

// Compute shortest distance between every single pair of vertices in a graph.
// Given adjacency list

// edges are represented as array of [[vertex, distance]]
// edges = [
//	[[1, 10], [2, 12]],                 // vertex - 0
//  [[2, 1], [5, 4], [3, 8], [4, 5]],   // vertex - 1
//  [[3, 6]],                           // vertex - 2
//  [[4, 6]],                           // vertex - 3
//  [[5, 3]],                           // vertex - 4
//  [],                                 // vertex - 5
// ]
//

// Equivalent Distance table

// +-----+-----+-----+-----+-----+-----+-----+
// | D() |  0  |  1  |  2  |  3  |  4  |  5  |
// +-----+-----+-----+-----+-----+-----+-----+
// |   0 | inf | 10  | 12  | inf | inf | inf |
// |   1 | inf | inf | 1   | 8   | 5   | 4   |
// |   2 | inf | inf | inf | 6   | inf | inf |
// |   3 | inf | inf | inf | inf | 6   | inf |
// |   4 | inf | inf | inf | inf | inf | 3   |
// |   5 | inf | inf | inf | inf | inf | inf |
// +-----+-----+-----+-----+-----+-----+-----+

// The shortest path from vertex i to vertex j that uses only up to k
// intermediate nodes is the shortest path that either does not use vertex k at all, or consists
// of the merging of the two paths vertex i to vertex k and vertex k to vertex j.  This leads to
// the formula:
// D[k,i,j] = min( D[k-1,i,j], D[k-1,i,k] + D[k-1,k,j])
// Previous best Previous best to vertex k, then best from k to j

type edge struct {
	vertex   int
	distance int
}

func AllParisShortest(edges [][]edge) [][]int {
	numVertices := len(edges)
	d := make([][]int, numVertices)
	for i := range d {
		d[i] = make([]int, numVertices)
		for j := 0; j < numVertices; j++ {
			d[i][j] = math.MaxInt32
		}
	}

	for i, vertexEdges := range edges {
		for _, edge := range vertexEdges {
			d[i][edge.vertex] = edge.distance
		}
	}

	for k := 0; k < numVertices; k++ {
		for i := 0; i < numVertices; i++ {
			for j := 0; j < numVertices; j++ {
				if d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}

	return d
}
