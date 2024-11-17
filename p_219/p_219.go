package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/contains-duplicate-ii/description/
func main() {
	arr := []int{1, 2, 3, 1}
	res := containsNearbyDuplicate(arr, 3)
	fmt.Println(res)
}
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		val, ok := m[v]
		if ok {
			if i-val <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}
