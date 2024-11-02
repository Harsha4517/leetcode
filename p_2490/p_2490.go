package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/circular-sentence/
func main() {
	res := isCircularSentence("leetcode exercises sound delightful")
	fmt.Println(res)
}

func isCircularSentence(sentence string) bool {
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			if sentence[i-1] != sentence[i+1] {
				return false
			}
		}
	}
	return sentence[0] == sentence[len(sentence)-1]
}
