package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/description/
func main() {
	arr := [][]int{
		{0, 1},
		{1, 1},
	}
	res := maxEqualRowsAfterFlips(arr)
	fmt.Println(res)
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	m := map[string]int{}
	result := 0
	for _, a := range matrix {
		s := make([]byte, len(a))
		for i, n := range a {
			if n == a[0] {
				s[i] = '1'
			}
		}
		str := string(s)
		m[str]++
		if m[str] > result {
			result = m[str]
		}
	}
	return result
}
