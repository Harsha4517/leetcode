package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/move-zeroes/description
func main() {
	var arr = []int{0, 1, 0, 3, 12}
	res := moveZeroes(arr)
	fmt.Println(res)
}

func moveZeroes(nums []int) []int {
	p := 0
	for q := 0; q < len(nums); q++ {
		if nums[q] != 0 {
			temp := nums[q]
			nums[q] = nums[p]
			nums[p] = temp
			p++
		}
	}
	return nums
}
