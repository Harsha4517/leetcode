package main

import "fmt"

const M = 30

// Leetcode url: https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/description/
func main() {
	arr := []int{1, 2, 3}
	res := minimumSubarrayLength(arr, 2)
	fmt.Println(res)
}
func updateCnt(num int, cur *int, cnt []int) {
	for i := 0; i < M; i++ {
		if num&(1<<i) != 0 {
			cnt[i]++
			if cnt[i] == 1 {
				*cur += 1 << i
			}
		}
	}
}

func removeCnt(num int, cur *int, cnt []int) {
	for i := 0; i < M; i++ {
		if num&(1<<i) != 0 {
			cnt[i]--
			if cnt[i] == 0 {
				*cur -= 1 << i
			}
		}
	}
}

func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}

	n := len(nums)
	l := 0
	cur := 0
	cnt := make([]int, M)
	res := n + 1

	for r, num := range nums {
		updateCnt(num, &cur, cnt)
		if cur >= k {
			for l <= r && cur >= k {
				removeCnt(nums[l], &cur, cnt)
				l++
			}
			res = min(res, r-(l-1)+1)
		}
	}

	if res == n+1 {
		return -1
	}
	return res
}
