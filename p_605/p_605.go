package main

import (
	"fmt"
)

// Leetcode url: https://leetcode.com/problems/can-place-flowers/description/
func main() {
	var arr = []int{1, 0, 0, 0, 1}
	res := canPlaceFlowers(arr, 1)
	fmt.Println(res)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	possibleRes := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			emptyLeft := i == 0 || flowerbed[i-1] == 0
			emptyRight := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if emptyLeft && emptyRight {
				possibleRes += 1
				flowerbed[i] = 1
			}
			if possibleRes >= n {
				return true
			}
		}
	}
	return false
}
