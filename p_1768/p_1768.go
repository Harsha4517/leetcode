package main

import (
	"fmt"
	"strings"
)

// Leetcode url: https://leetcode.com/problems/merge-strings-alternately/description/
func main() {
	var res string = mergeAlternately("abcd", "pqrstu")
	fmt.Println(res)
}

func mergeAlternately(string1 string, string2 string) string {
	var mergeRes strings.Builder
	len1 := len(string1)
	len2 := len(string2)

	minValue := len1
	if len2 < len1 {
		minValue = len2
	}

	for i := 0; i < minValue; i++ {
		mergeRes.WriteString(string(string1[i]))
		mergeRes.WriteString(string(string2[i]))
	}

	for i := minValue; i < len1; i++ {
		mergeRes.WriteString(string(string1[i]))
	}
	for i := minValue; i < len2; i++ {
		mergeRes.WriteString(string(string2[i]))
	}
	return mergeRes.String()
}
