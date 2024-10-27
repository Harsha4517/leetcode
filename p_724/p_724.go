package main

import (
	"fmt"
)

// Leetcode Url: https://leetcode.com/problems/find-pivot-index/description/
func main() {
	var arr = []int{1, 7, 3, 6, 5, 6}
	res := pivotIndex(arr)
	fmt.Println(res)
}

func pivotIndex(nums []int) int {
	sum := 0
	prefixSum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if prefixSum == sum-prefixSum-nums[i] {
			return i
		}
		prefixSum += nums[i]
	}
	return -1
}
