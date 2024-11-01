package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/delete-characters-to-make-fancy-string/
func main() {
	res := makeFancyString(("leeetcode"))
	fmt.Println(res)
}

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	result := []byte{}
	temp := 0
	for i := 0; i < len(s); i++ {
		if i < 2 || !(s[i] == result[temp-1] && s[i] == result[temp-2]) {
			result = append(result, s[i])
			temp = temp + 1
		}
	}
	return string(result)
}
