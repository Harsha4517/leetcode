package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/container-with-most-water/description/
func main() {
	var arr = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(arr)
	fmt.Println(res)
}

func maxArea(height []int) int {
	maximum_water := 0
	p := 0
	q := len(height) - 1
	temp := 0
	for p < q {
		if height[p] < height[q] {
			temp = height[p] * (q - p)
			p++
		} else {
			temp = height[q] * (q - p)
			q--
		}
		maximum_water = max(temp, maximum_water)

	}
	return maximum_water
}
