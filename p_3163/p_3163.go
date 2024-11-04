package main

import (
	"fmt"
	"strconv"
)

// Leetcode url: https://leetcode.com/problems/string-compression-iii/description/
func main() {
	res := compressedString("aaaaaaaaaaaaaabb")
	fmt.Println(res)
}

func compressedString(word string) string {
	ans := []byte{}
	count := 1

	for i := 1; i <= len(word); i++ {
		if i == len(word) || word[i] != word[i-1] || count == 9 {
			ans = append(ans, []byte(strconv.Itoa(count))...)
			ans = append(ans, word[i-1])
			count = 1
		} else {
			count++
		}
	}

	return string(ans)
}
