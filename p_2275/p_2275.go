package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/largest-combination-with-bitwise-and-greater-than-zero/description/
func main() {
	var arr = []int{16, 17, 71, 62, 12, 24, 14}
	res := largestCombination(arr)
	fmt.Println(res)
}

func largestCombination(candidates []int) int {
	max_count := 0
	for i := 0; i < 24; i++ {
		count := 0
		for j := 0; j < len(candidates); j++ {
			if (candidates[j] & (1 << i)) != 0 {
				count += 1
			}
		}
		max_count = max(max_count, count)
	}
	return max_count
}
