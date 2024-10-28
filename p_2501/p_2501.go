package main

import (
	"fmt"
	"math"
	"sort"
)

// Leetcode Url: https://leetcode.com/problems/longest-square-streak-in-an-array/description/
func main() {
	var arr = []int{4, 3, 6, 16, 8, 2}
	res := longestSquareStreak(arr)
	fmt.Println(res)
}

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	streak := make(map[int]int)
	maxStreak := 0
	for i := 0; i < len(nums); i++ {
		rootNum := int(math.Sqrt(float64(nums[i])))
		if rootNum*rootNum == nums[i] && streak[rootNum] != 0 {
			streak[nums[i]] = streak[rootNum] + 1
		} else {
			streak[nums[i]] = 1
		}

		maxStreak = max(maxStreak, streak[nums[i]])
	}

	if maxStreak < 2 {
		return -1
	}
	return maxStreak
}
