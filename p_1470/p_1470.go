package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/shuffle-the-array/description/
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	res := shuffle(arr, 3)
	fmt.Println(res)
}
func shuffle(nums []int, n int) []int {
	var newarr []int
	for i := 0; i < n; i++ {
		newarr = append(newarr, nums[i], nums[i+n])
	}
	return newarr
}
