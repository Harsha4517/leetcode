package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/description/
func main() {
	res := takeCharacters("aabaaaacaabc", 2)
	fmt.Println(res)
}

func takeCharacters(s string, k int) int {
	count_array := []int{0, 0, 0}
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			count_array[0] += 1
		} else if s[i] == 'b' {
			count_array[1] += 1
		} else {
			count_array[2] += 1
		}
	}
	for i := 0; i < 3; i++ {
		if count_array[i] < k {
			return -1
		}
	}
	left, ans := 0, 0
	window := []int{0, 0, 0}
	for right := 0; right < n; right++ {
		window[s[right]-'a'] += 1
		for left <= right && (count_array[0]-window[0] < k || count_array[1]-window[1] < k || count_array[2]-window[2] < k) {
			window[s[left]-'a'] -= 1
			left += 1
		}
		ans = max(ans, right-left+1)
	}
	return n - ans
}
