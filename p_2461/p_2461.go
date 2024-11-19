package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/
func main() {
	arr := []int{1, 5, 4, 2, 9, 9, 9}
	res := maximumSubarraySum(arr, 3)
	fmt.Println(res)
}

func maximumSubarraySum(nums []int, k int) int64 {
	freqMap := map[int]int{}
	sum, ans := 0, 0

	for i := range k {
		cur := nums[i]
		sum += cur
		freqMap[cur]++
	}
	if len(freqMap) == k {
		ans = max(ans, sum)
	}

	for i := k; i < len(nums); i++ {
		cur, pre := nums[i], nums[i-k]
		sum += cur - pre
		if freqMap[pre] == 1 {
			delete(freqMap, pre)
		} else {
			freqMap[pre]--
		}
		freqMap[cur]++
		if len(freqMap) == k {
			ans = max(ans, sum)
		}
	}
	return int64(ans)
}
