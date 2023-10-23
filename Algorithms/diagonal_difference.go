package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int) int {
	sum := 0

	for i, j := 0, len(arr[0])-1; j >= 0; i, j = i+1, j-1 {
		sum += arr[i][i] - arr[i][j]
	}
	return int(math.Abs(float64(sum)))
}

func main() {
	// Example matrix
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	result := diagonalDifference(matrix)
	fmt.Printf("Absolute diagonal difference: %d\n", result)
}
