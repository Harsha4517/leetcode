package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Leetcode url: https://leetcode.com/problems/merge-two-sorted-lists/description/
func main() {
	var arr1 = &ListNode{1, nil}
	var arr2 = &ListNode{1, nil}
	arr1.Next, arr2.Next = &ListNode{2, nil}, &ListNode{3, nil}
	arr1.Next.Next, arr2.Next.Next = &ListNode{4, nil}, &ListNode{4, nil}
	res := mergeTwoLists(arr1, arr2)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return result.Next
}
