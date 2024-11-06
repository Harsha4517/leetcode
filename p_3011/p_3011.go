package main

import (
	"fmt"
	"math"
	"math/bits"
)

// Leetcode url: https://leetcode.com/problems/find-if-array-can-be-sorted/description/
func main() {
	var arr = []int{8, 4, 2, 30, 15}
	res := canSortArray(arr)
	fmt.Println(res)
}

func canSortArray(nums []int) bool {
	max_of_segment := nums[0]
	min_of_segment := nums[0]
	max_of_previous_segment := math.MinInt
	for i := 1; i < len(nums); i++ {
		if numberOfBit(nums[i]) == numberOfBit(nums[i-1]) {
			max_of_segment = max(max_of_segment, nums[i])
			min_of_segment = min(min_of_segment, nums[i])
		} else {
			if min_of_segment < max_of_previous_segment {
				return false
			}
			max_of_previous_segment = max_of_segment
			max_of_segment = nums[i]
			min_of_segment = nums[i]
		}
	}
	if min_of_segment < max_of_previous_segment {
		return false
	}
	return true
}

func numberOfBit(num int) int {
	return int(bits.OnesCount(uint(num)))
}
