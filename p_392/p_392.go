package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/is-subsequence/description/
func main() {
	var res bool = isSubsequence("abc", "ahbgdc")
	fmt.Println(res)
}

func isSubsequence(s string, t string) bool {
	k := 0
	for i := 0; i < len(t); i++ {
		if k == len(s) {
			return true
		}
		if s[k] == t[i] {
			k++
		}
	}
	return k == len(s)
}
