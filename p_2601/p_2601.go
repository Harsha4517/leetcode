package main

import (
	"fmt"
	"math"
)

// Leetcode url: https://leetcode.com/problems/prime-subtraction-operation/description/
func main() {
	arr := []int{4, 9, 6, 10}
	res := primeSubOperation(arr)
	fmt.Println(res)
}

func primeSubOperation(nums []int) bool {
	maxElement := nums[0]
	for i := 1; i < len(nums); i++ {
		maxElement = max(maxElement, nums[i])
	}
	sieve := make([]int, maxElement+1)
	for i := 2; i < maxElement+1; i++ {
		sieve[i] = 1
	}

	for i := 2; i <= int(math.Sqrt(float64(maxElement+1))); i++ {
		if sieve[i] == 1 {
			for j := i * i; j <= maxElement; j += i {
				sieve[j] = 0
			}
		}
	}
	currValue := 1
	i := 0
	for i < len(nums) {
		difference := nums[i] - currValue
		if difference < 0 {
			return false
		}
		if sieve[difference] == 1 || difference == 0 {
			i++
			currValue++
		} else {
			currValue++
		}
	}
	return true
}
