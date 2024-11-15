package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/description/
func main() {
	arr := []int{1, 2, 3, 10, 4, 2, 3, 5}
	res := findLengthOfShortestSubarray(arr)
	fmt.Println(res)
}

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	left := 0
	for left < n-1 && arr[left] <= arr[left+1] {
		left++
	}
	if left == n-1 {
		return 0
	}
	right := n - 1
	for right >= 0 && arr[right] >= arr[right-1] {
		right--
	}
	result := min(n-left-1, right)
	i := 0
	j := right
	for i <= left && j < n {
		if arr[i] <= arr[j] {
			result = min(j-i-1, result)
			i++
		} else {
			j++
		}
	}
	return result
}
