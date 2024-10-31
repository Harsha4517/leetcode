package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
func main() {
	res := strStr("sadbutsad", "sad")
	fmt.Println(res)
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	h := len(haystack)
	n := len(needle)
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
