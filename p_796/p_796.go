package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/rotate-string/description/
func main() {
	res := rotateString("abcde", "cdeab")
	fmt.Println(res)
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	temp := s + s
	n := len(temp)
	for i := 0; i < n-len(goal); i++ {
		current := temp[i : i+len(goal)]
		if current == goal {
			return true
		}
	}
	return false
}
