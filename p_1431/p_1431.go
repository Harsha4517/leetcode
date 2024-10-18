package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description
func main() {
	var arr = []int{1, 2, 3, 4, 5}
	res := kidsWithCandies(arr, 3)
	fmt.Println(res)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	totalKids := len(candies)
	res := make([]bool, totalKids)
	maxValue := candies[0]
	for i := 1; i < totalKids; i++ {
		if candies[i] > maxValue {
			maxValue = candies[i]
		}
	}
	for i := 0; i < totalKids; i++ {
		if candies[i]+extraCandies >= maxValue {
			res[i] = true
		}
	}
	return res
}
