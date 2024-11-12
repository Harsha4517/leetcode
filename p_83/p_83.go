package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Leetcode url: https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
func main() {
	var arr1 = &ListNode{1, nil}
	arr1.Next = &ListNode{1, nil}
	arr1.Next.Next = &ListNode{2, nil}
	res := deleteDuplicates(arr1)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
