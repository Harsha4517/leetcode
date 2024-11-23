package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/reverse-bits/description/
func main() {
	res := reverseBits(0b00000010100101000001111010011100)
	fmt.Println(res)
}

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 + num&1
		num >>= 1
	}
	return res
}
