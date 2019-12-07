package list

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func NewList(nums []int) *Node {
	root := &Node{Val: nums[0]}
	prev := root
	for i := 1; i < len(nums); i++ {
		prev.Next = &Node{Val: nums[i]}
		prev = prev.Next
	}
	return root
}

func PrintList(h *Node) {
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

func Equal(h1, h2 *Node) bool {
	for h1 != nil && h2 != nil {
		if h1.Val != h2.Val {
			return false
		}
		h1, h2 = h1.Next, h2.Next
	}

	if h1 != nil {
		return false
	}
	if h2 != nil {
		return false
	}
	return true
}

func Reverse(h *Node) *Node {
	var prev *Node
	for h != nil {
		h.Next, h, prev = prev, h.Next, h
	}
	return prev
}

func Add(h1, h2 *Node) *Node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	h1 = Reverse(h1)
	h2 = Reverse(h2)

	head := &Node{}
	prev := head
	r, sum := 0, 0
	for h1 != nil || h2 != nil || r != 0 {
		if h1 != nil {
			sum, h1 = sum+h1.Val, h1.Next

		}
		if h2 != nil {
			sum, h2 = sum+h2.Val, h2.Next
		}
		sum += r
		r = sum / 10

		prev.Next = &Node{Val: sum % 10}
		prev, sum = prev.Next, 0
	}

	return Reverse(head.Next)
}
