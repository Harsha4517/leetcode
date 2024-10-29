package main

import (
	"fmt"
)

// Leetcode Url: https://leetcode.com/problems/longest-common-prefix/description/
func main() {
	var arr = []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(arr)
	fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}