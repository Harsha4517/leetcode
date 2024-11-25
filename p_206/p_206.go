package main

import "fmt"

// Leetcode url: https://leetcode.com/problems/reverse-linked-list/description/
func main() {
	var arr1 = &ListNode{1, nil}
	arr1.Next = &ListNode{2, nil}
	arr1.Next.Next = &ListNode{3, nil}
	res := reverseList(arr1)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
