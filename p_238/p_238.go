package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/product-of-array-except-self/description/
func main() {
	var arr = []int{1, 2, 3, 4}
	res := productExceptSelf(arr)
	fmt.Println(res)
}

func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	res := make([]int, numsLen)
	prefix := 1
	for i := 0; i < numsLen; i++ {
		res[i] = prefix
		prefix *= nums[i]
	}
	suffix := 1
	for i := numsLen - 1; i > -1; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res
}
