package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/description/
func main() {
	arr := []int{1, 2, 3, 4, 3, 2, 5}
	res := resultsArray(arr, 3)
	fmt.Println(res)
}

func resultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	arr_len := len(nums)
	results := make([]int, arr_len-k+1)

	for i := range results {
		results[i] = -1
	}
	cons_count := 1
	for i := 0; i < arr_len-1; i++ {
		if nums[i]+1 == nums[i+1] {
			cons_count += 1
		} else {
			cons_count = 1
		}
		if cons_count >= k {
			results[i-k+2] = nums[i+1]
		}
	}
	return results
}
