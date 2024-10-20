package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/greatest-common-divisor-of-strings/description/
func main() {
	res := gcdOfStrings("ABCABCABC", "ABC")
	fmt.Println(res)
}

func gcdOfStrings(s1 string, s2 string) string {
	if s1+s2 != s2+s1 {
		return ""
	}
	x := gcd(len(s1), len(s2))
	return s1[:x]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
