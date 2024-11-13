package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/score-of-a-string/description/
func main() {
	res := scoreOfString("hello")
	fmt.Println(res)
}
func scoreOfString(s string) int {
	var res int = 0
	for i := 0; i < len(s)-1; i++ {
		res += getDiff(int(s[i]), int(s[i+1]))
	}
	return res
}

func getDiff(n1 int, n2 int) int {
	if n1 > n2 {
		return n1 - n2
	}
	return n2 - n1
}
