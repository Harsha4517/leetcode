package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/find-the-highest-altitude/description/
func main() {
	var arr = []int{-5, 1, 5, 0, -7}
	res := largestAltitude(arr)
	fmt.Println(res)
}

func largestAltitude(gain []int) int {
	res := 0
	temp := 0
	for i := 0; i < len(gain); i++ {
		temp += gain[i]
		res = max(res, temp)
	}
	return res
}
