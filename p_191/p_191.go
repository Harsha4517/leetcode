package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/number-of-1-bits/description/
func main() {
	res := hammingWeight(11)
	fmt.Println(res)
}

func hammingWeight(n int) int {
	res := 0
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
