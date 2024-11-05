package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/minimum-number-of-changes-to-make-binary-string-beautiful/description/
func main() {
	res := minChanges("1001")
	fmt.Println(res)
}

func minChanges(s string) int {
	res := 0
	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			res += 1
		}
	}
	return res
}
