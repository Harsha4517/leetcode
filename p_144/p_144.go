package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Leetcode url: https://leetcode.com/problems/binary-tree-preorder-traversal/
func main() {
	var arr1 = &TreeNode{1, nil, nil}
	var arr2 = &TreeNode{2, nil, nil}
	arr1.Right = arr2
	var arr3 = &TreeNode{3, nil, nil}
	arr2.Left = arr3
	res := preorderTraversal(arr1)
	fmt.Println(res)
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}
