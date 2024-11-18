package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/defuse-the-bomb/description/
func main() {
	arr := []int{5, 7, 1, 4}
	res := decrypt(arr, 3)
	fmt.Println(res)
}

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))
	if k == 0 {
		return result
	}
	start, end, sum := 1, k, 0
	if k < 0 {
		start = len(code) + k
		end = len(code) - 1
	}
	for i := start; i < end+1; i++ {
		sum += code[i]
	}
	for i := 0; i < len(code); i++ {
		result[i] = sum
		sum -= code[start%len(code)]
		sum += code[(end+1)%len(code)]
		start += 1
		end += 1
	}
	return result
}
