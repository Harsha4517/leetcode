package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/maximum-xor-for-each-query/description/
func main() {
	var arr = []int{0, 1, 1, 3}
	res := getMaximumXor(arr, 2)
	fmt.Println(res)
}

func getMaximumXor(nums []int, maximumBit int) []int {
	n, val, k := len(nums), 0, (1<<maximumBit)-1
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		val ^= nums[i]
		ans[n-i-1] = val ^ k
	}
	return ans
}
