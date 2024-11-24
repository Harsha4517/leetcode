package main

import (
	"fmt"
	"math"
)

// Leetcode url: https://leetcode.com/problems/maximum-matrix-sum/description/
func main() {
	arr := [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}
	res := maxMatrixSum(arr)
	fmt.Println(res)
}

func maxMatrixSum(matrix [][]int) int64 {
	totalCount := 0
	totalNCount := 0
	minVal := math.MaxInt
	for i, row := range matrix {
		for j := 0; j < len(row); j++ {
			totalCount += absInt(matrix[i][j])
			minVal = min(minVal, absInt(matrix[i][j]))
			if matrix[i][j] < 0 {
				totalNCount += 1
			}
		}
	}
	if totalNCount&1 != 0 {
		return int64(totalCount - 2*minVal)
	}
	return int64(totalCount)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
