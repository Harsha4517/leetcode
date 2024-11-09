package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/minimum-array-end/description/
func main() {
	res := minEnd(3, 4)
	fmt.Println(res)
}

func minEnd(n int, x int) int64 {
	result := x
	n -= 1
	mask := 1
	for n > 0 {
		if (mask & x) == 0 {
			result |= (n & 1) * mask
			n >>= 1
		}
		mask <<= 1
	}
	return int64(result)
}
